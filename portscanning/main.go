package main
import (
	"fmt"
	"net"
	"sync"
	"time"
	"os/exec"
	"bufio"
	"os"
	"strings"
)

func pingHost(ip string) bool{
	// ping -c 1 send 1 packet
	// -W 1 timeout
	cmd := exec.Command("ping","-c","1","-W","1", ip)
	err := cmd.Run()
	return err == nil
}

func getMac(ip string) string{
	file, err := os.Open("/proc/net/arp")
	if err != nil {
		return ""
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) >= 4 && fields[0] == ip {
			return fields[3]
		}
	}
	return ""
}

func scanHost(wg *sync.WaitGroup, host string){
	defer wg.Done()
	mac := getMac(host)
	fmt.Printf("Alive: %s MAC:%s\n", host, mac)
}

func scanPort(wg *sync.WaitGroup, host string, port int){
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)

	if err != nil{
		return
	}
	conn.Close()
	fmt.Println("[OPEN] Port", port, "Host[",host,"]")
}

func main(){
	base := "192.168.1."
	var wg sync.WaitGroup

	for i:=0; i<=254; i++{
		host := fmt.Sprintf("%s%d", base, i)
		wg.Add(1)
		go func(host string){
			defer wg.Done()
			if !pingHost(host){
				return
			}
			mac := getMac(host)
			fmt.Printf("Alive: %s MAC: %s\n", host, mac)
			for p:=0; p<=1024; p++{
				scanPort(&wg, host, p) 
			}
		}(host)
	}

	wg.Wait()
}