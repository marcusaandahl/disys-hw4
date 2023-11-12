package main

import (
	"bufio"
	"context"
	"fmt"
	gRPC "github.com/marcusaandahl/disys-hw4/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

// Client id, port and whether it has a connected peer
var id int32
var port int32 = 4000
var hasNextConnection = false

// Client next peer in token ring credentials
var nextId int32
var nextPort int32
var nextPeer gRPC.PeerProtoClient
var nextPeerConn *grpc.ClientConn

// PeerAsServer server declaration for Client as peer server
type PeerAsServer struct {
	gRPC.UnimplementedPeerProtoServer
	mutex sync.Mutex
}

// Client status
var wantsToAccess = false
var isAccessingResource = false

func main() {
	// EXIT FALLBACK - unstable
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTSTP, syscall.SIGKILL, syscall.SIGTERM)
	go func() {
		<-signals
		exit()
	}()

	// Get running ports in project
	var ports = readPorts()

	// Wait for server spin up
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		startServer(&wg)
	}()
	wg.Wait()

	// Connect to network if any other peers in network
	if len(ports) > 0 {
		selectedPort := ports[len(ports)-1]

		connectToPort(selectedPort)
	}

	// Keep alive scanner with commands
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		switch strings.ToLower(command) {
		case "access":
			fmt.Println("Attempting to access resource, please wait...")
			startVote()
			break
		case "release":
			releaseResource()
			break
		case "exit":
			exit()
			break
		default:
			fmt.Println("Wrong command, try 'access', 'release' or 'exit'")
		}
	}
}

// Starts a vote for accessing a resource
func startVote() {
	if isAccessingResource {
		// If client already has access to resource - continue
		fmt.Println("You already have access to the critical resource, write 'release' to release it")
	} else if !hasNextConnection {
		// If client is alone in network - give access
		isAccessingResource = true
		fmt.Println("You now have access to resource, write 'release' to release resource")
	} else {
		// If client is not alone in network and doesn't have resource - start a vote for it
		wantsToAccess = true
		_, _ = nextPeer.Vote(context.Background(), &gRPC.VoteRequest{
			RequesterId: id,
			DeniedBy:    -1,
			Denied:      false,
		})
	}
}

// Client releases resource
func releaseResource() {
	if isAccessingResource {
		isAccessingResource = false
		fmt.Println("You no longer have access to the resource")
	} else {
		fmt.Println("You do not have access to the resource")
	}
}

// Connects to peer network
func connectToPort(tempPort string) {
	// Requests its credentials from network
	tempOpts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	tempConn, err := grpc.Dial(fmt.Sprintf(":%v", tempPort), tempOpts...)
	check(err)
	tempNextPeer := gRPC.NewPeerProtoClient(tempConn)
	tempNextPeerConn := tempConn
	res, err := tempNextPeer.CheckConnection(context.Background(), &gRPC.ConnectRequest{
		NewPeerPort: port,
	})
	check(err)
	err = tempNextPeerConn.Close()
	check(err)

	// Sets itself up as a new client in the network given the received credentials
	id = res.GetNewPeerId()
	nextId = res.GetNewPeerNextPeerId()
	nextPort = res.GetNewPeerNextPeerPort()

	fmt.Printf("Client on port %v has ID %v - Next Client on port %v has ID %v \n \n", port, id, nextPort, nextId)

	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(fmt.Sprintf(":%v", nextPort), opts...)
	check(err)

	nextPeer = gRPC.NewPeerProtoClient(conn)
	nextPeerConn = conn

	hasNextConnection = true
}

// Vote handles vote in peer token ring
func (s *PeerAsServer) Vote(context context.Context, voteRequest *gRPC.VoteRequest) (*emptypb.Empty, error) {
	if voteRequest.RequesterId == id {
		// Client initiated request as gone around
		//fmt.Printf("My request has gone around and is %v by %v \n", voteRequest.GetDenied(), voteRequest.GetDeniedBy())
		if voteRequest.GetDenied() {
			// Request was denied
			//fmt.Printf("Request from Client %v denied by client %v \n", id, voteRequest.GetDeniedBy())
			go func() {
				// Retry
				startVote()
			}()
		} else {
			// Was not denied - gets access to resource
			isAccessingResource = true
			wantsToAccess = false
			fmt.Println("You now have access to resource, write 'release' to release resource")
		}
		return &emptypb.Empty{}, nil
	} else {
		// Request from other peer going through
		if voteRequest.GetDenied() {
			// Request is denied - send to next peer
			//fmt.Printf("Sending request from %v which is %v \n", voteRequest.GetRequesterId(), voteRequest.GetDenied())
			_, _ = nextPeer.Vote(context, voteRequest)
			return &emptypb.Empty{}, nil
		}

		if wantsToAccess || isAccessingResource {
			// If client wants to access or has access
			if id > voteRequest.GetRequesterId() || isAccessingResource {
				// Deny request
				//fmt.Printf("Denied request from %v \n", voteRequest.GetRequesterId())
				_, _ = nextPeer.Vote(context, &gRPC.VoteRequest{
					RequesterId: voteRequest.GetRequesterId(),
					DeniedBy:    id,
					Denied:      true,
				})
				return &emptypb.Empty{}, nil
			} else {
				// Client can't deny request
				//fmt.Printf("I can't deny request from %v \n", voteRequest.GetRequesterId())
				_, _ = nextPeer.Vote(context, voteRequest)
				return &emptypb.Empty{}, nil
			}
		} else {
			// Client doesn't want access - send request in peer token ring
			//fmt.Printf("I don't want access %v \n", voteRequest.GetRequesterId())
			_, _ = nextPeer.Vote(context, voteRequest)
			return &emptypb.Empty{}, nil
		}
	}
}

// CheckConnection ensures the new client gets its credentials, whilst current peers in token ring make space for it
func (s *PeerAsServer) CheckConnection(context context.Context, connectRequest *gRPC.ConnectRequest) (*gRPC.ConnectResponse, error) {
	if !hasNextConnection {
		// 1 peer in the chain - adding new peer
		id = 0
		nextId = id + 1
		nextPort = connectRequest.GetNewPeerPort()

		//fmt.Printf("I'm the only peer in chain - adding peer %v at port %v \n", nextId, nextPort)

		opts := []grpc.DialOption{
			grpc.WithBlock(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}

		conn, err := grpc.Dial(fmt.Sprintf(":%v", nextPort), opts...)
		check(err)

		nextPeer = gRPC.NewPeerProtoClient(conn)
		nextPeerConn = conn
		hasNextConnection = true

		// What id should i have
		// Where do i point

		return &gRPC.ConnectResponse{
			NewPeerNextPeerId:   id,
			NewPeerNextPeerPort: port,
			NewPeerId:           nextId,
		}, nil
	} else if nextId < id {
		// Last in chain - adding new peer
		var newPeerResponse = gRPC.ConnectResponse{
			NewPeerNextPeerId:   nextId,
			NewPeerNextPeerPort: nextPort,
			NewPeerId:           id + 1,
		}

		//fmt.Printf("I'm the last peer in chain - adding peer %v at port %v pointing to peer %v on port %v \n", id+1, connectRequest.GetNewPeerPort(), nextId, nextPort)

		nextId = id + 1
		nextPort = connectRequest.GetNewPeerPort()

		opts := []grpc.DialOption{
			grpc.WithBlock(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}

		conn, err := grpc.Dial(fmt.Sprintf(":%v", nextPort), opts...)
		check(err)

		nextPeer = gRPC.NewPeerProtoClient(conn)
		nextPeerConn = conn
		hasNextConnection = true

		return &newPeerResponse, nil
	} else {
		// Not last in chain - sending request on
		res, err := nextPeer.CheckConnection(context, connectRequest)
		return res, err
	}
}

// LeaveNetwork ensures peer safely leaves network, whilst other peers reconnect the token ring
func (s *PeerAsServer) LeaveNetwork(context context.Context, leaveRequest *gRPC.LeaveRequest) (*emptypb.Empty, error) {
	//Determine if right recipient -> if not send to next one
	if nextId == leaveRequest.GetTraitorId() {
		//The leaving node is my successor
		nextId = leaveRequest.GetTraitorNextId()
		nextPort = leaveRequest.GetTraitorNextPort()

		//fmt.Printf("Client on port %v has ID %v - Next Client on port %v has ID %v \n \n", port, id, nextPort, nextId)

		opts := []grpc.DialOption{
			grpc.WithBlock(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}

		conn, err := grpc.Dial(fmt.Sprintf(":%v", nextPort), opts...)
		check(err)

		nextPeer = gRPC.NewPeerProtoClient(conn)
		nextPeerConn = conn

		//Logging purposes
		//fmt.Printf("Received leave from: %v. My new successor: %v at port %v \n", leaveRequest.GetTraitorId(), nextId, nextPort)

		return &emptypb.Empty{}, nil
	} else {
		//The leaving node is NOT my successor -> relay message
		res, err := nextPeer.LeaveNetwork(context, leaveRequest)

		//Logging purposes
		//fmt.Printf("Passing along leave request from %v to: %v at port %v \n", leaveRequest.GetTraitorId(), nextId, nextPort)

		return res, err
	}
}

// Starts the client as server, such that clients can access it
func startServer(wg *sync.WaitGroup) {
	var connString = fmt.Sprintf("localhost:%v", port)

	list, err := net.Listen("tcp", connString)

	var attempts = 1
	for err != nil && attempts < 40 {
		attempts++
		port++
		var connString = fmt.Sprintf("localhost:%v", port)
		list, err = net.Listen("tcp", connString)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	server := &PeerAsServer{}

	gRPC.RegisterPeerProtoServer(grpcServer, server)

	log.Printf("Client listening at %v\n", list.Addr())

	addPortToFile()

	wg.Done()

	if err := grpcServer.Serve(list); err != nil {
		removePortFromFile()
		log.Fatalf("failed to serve %v", err)
	}
}

// Reads ports in usedPorts file
func readPorts() []string {
	dat, err := os.Open("./usedPorts.txt")
	check(err)

	scan := bufio.NewScanner(dat)

	var ports []string

	for scan.Scan() {
		var text = scan.Text()
		if len(text) > 0 {
			ports = append(ports, text)
		}
	}

	defer closeFile(dat)

	return ports
}

// Adds new port to usedPorts file
func addPortToFile() {
	var ports = readPorts()
	if !contains(ports, fmt.Sprintf("%v", port)) {
		ports = append(ports, fmt.Sprintf("%v", port))
		fileData := strings.Join(ports, "\n")

		overwriteFile("./usedPorts.txt", fileData)
	}
}

// Removes port from usedPorts file
func removePortFromFile() {
	var ports = readPorts()
	if contains(ports, fmt.Sprintf("%v", port)) {
		ports = removeElementFromSlice(ports, fmt.Sprintf("%v", port))
	}

	content := strings.Join(ports, "\n")

	overwriteFile("./usedPorts.txt", content)
}

// Overrides file
func overwriteFile(path string, content string) {
	err := os.WriteFile(path, []byte(content), 0644)
	check(err)
}

// Contains function as it is not built-in in golang 🤡
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Removes element from slice
func removeElementFromSlice(s []string, e string) []string {
	var newSlice []string

	for _, a := range s {
		if a != e {
			newSlice = append(newSlice, a)
		}
	}

	return newSlice
}

// Closes file
func closeFile(dat *os.File) {
	err := dat.Close()
	check(err)
}

// Exits process
func exit() {
	// Releases any potential resource
	wantsToAccess = false
	isAccessingResource = false

	// Disconnects if more other peers in token ring
	if hasNextConnection {
		_, err := nextPeer.LeaveNetwork(context.Background(), &gRPC.LeaveRequest{
			TraitorId:       id,
			TraitorNextId:   nextId,
			TraitorNextPort: nextPort,
		})
		check(err)
	}

	// Removes its port from port file whilst closing its connection
	removePortFromFile()
	if hasNextConnection {
		err := nextPeerConn.Close()
		check(err)
	}

	os.Exit(0)
}

// Error panicking function
func check(e error) {
	if e != nil {
		panic(e)
	}
}
