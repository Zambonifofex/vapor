// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/ufoot/vapor/go/vpp2papi"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  HostStatus Status()")
	fmt.Fprintln(os.Stderr, "  GetSuccessorsResponse GetSuccessors(GetSuccessorsRequest request)")
	fmt.Fprintln(os.Stderr, "  LookupResponse Lookup(LookupRequest request)")
	fmt.Fprintln(os.Stderr, "  void ping()")
	fmt.Fprintln(os.Stderr, "  Version getVersion()")
	fmt.Fprintln(os.Stderr, "  Package getPackage()")
	fmt.Fprintln(os.Stderr, "  i64 uptime()")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := vpp2papi.NewVpP2pApiClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "Status":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "Status requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.Status())
		fmt.Print("\n")
		break
	case "GetSuccessors":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetSuccessors requires 1 args")
			flag.Usage()
		}
		arg16 := flag.Arg(1)
		mbTrans17 := thrift.NewTMemoryBufferLen(len(arg16))
		defer mbTrans17.Close()
		_, err18 := mbTrans17.WriteString(arg16)
		if err18 != nil {
			Usage()
			return
		}
		factory19 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt20 := factory19.GetProtocol(mbTrans17)
		argvalue0 := vpp2papi.NewGetSuccessorsRequest()
		err21 := argvalue0.Read(jsProt20)
		if err21 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetSuccessors(value0))
		fmt.Print("\n")
		break
	case "Lookup":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Lookup requires 1 args")
			flag.Usage()
		}
		arg22 := flag.Arg(1)
		mbTrans23 := thrift.NewTMemoryBufferLen(len(arg22))
		defer mbTrans23.Close()
		_, err24 := mbTrans23.WriteString(arg22)
		if err24 != nil {
			Usage()
			return
		}
		factory25 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt26 := factory25.GetProtocol(mbTrans23)
		argvalue0 := vpp2papi.NewLookupRequest()
		err27 := argvalue0.Read(jsProt26)
		if err27 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Lookup(value0))
		fmt.Print("\n")
		break
	case "ping":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "Ping requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.Ping())
		fmt.Print("\n")
		break
	case "getVersion":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetVersion requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetVersion())
		fmt.Print("\n")
		break
	case "getPackage":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetPackage requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetPackage())
		fmt.Print("\n")
		break
	case "uptime":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "Uptime requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.Uptime())
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
