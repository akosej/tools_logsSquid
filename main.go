package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"net"
	"os"
	"regexp"
	"strings"
)

func main() {
	nameFile := flag.String("f", "access.log", "file path to read from")
	flag.Parse()
	fmt.Println("Extract information from the " + *nameFile + " file")
	file, err := os.Open(*nameFile)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Check if the file already exists
	_, err = os.Stat(*nameFile + ".copy")
	if os.IsNotExist(err) {
		// El archivo no existe, crearlo
		file, err := os.Create(*nameFile + ".copy")
		if err != nil {
			fmt.Println("Error al crear el archivo:", err)
			return
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)
		fmt.Println("Se ha creado correctamente el archivo.")
	}
	archivo, _ := os.OpenFile(*nameFile+".copy", os.O_WRONLY|os.O_TRUNC, 0666)
	// Clear file content
	if err := archivo.Truncate(0); err != nil {
		fmt.Println("Error al borrar el contenido del archivo:", err)
		return
	}

	// Read Squid access log file
	data, err := os.ReadFile(*nameFile)
	if err != nil {
		panic(err)
	}
	// Regular expression to detect IP addresses
	ipRegex := regexp.MustCompile(`(?:\d{1,3}\.){3}\d{1,3}`)
	// Regular expression to exclude the range 10.26.0.0/16
	negativeRegex := regexp.MustCompile(`^10\.26\.`)

	ips := ipRegex.FindAllString(string(data), -1)
	// create and start new bar
	bar := pb.Full.Start(len(ips))
	var filteredIPs []string
	var filteredIPsNoDomain []string
	storage := map[string]string{}
	for _, ip := range ips {
		bar.Increment()
		// Exclude ip within the range 10.26.0.0/16
		if !negativeRegex.MatchString(ip) {
			// Check if the "ip" value exists within the saved slices
			if !strings.Contains(strings.Join(filteredIPs, ","), ip) && !strings.Contains(strings.Join(filteredIPsNoDomain, ","), ip) {
				domain := getDomainName(ip)
				if domain != "" {
					filteredIPs = append(filteredIPs, ip)
					storage[ip] = domain
				} else {
					filteredIPsNoDomain = append(filteredIPsNoDomain, ip)
				}
			}
		}
	}
	// finish bar
	bar.Finish()
	fmt.Println("Create a copy of the " + *nameFile + ".copy file with the domains found")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, ip := range filteredIPs {
			line = strings.Replace(line, ip, storage[ip], -1)
		}
		_ = AppendStrFile(*nameFile+".copy", line)
		_ = AppendStrFile(*nameFile+".copy", "\n")
	}
}

// AppendStrFile -- Add line at the end of the file
func AppendStrFile(path, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}

// Function to look up the domain name corresponding to an IP address
func getDomainName(ipAddress string) string {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		//fmt.Println("Dirección IP inválida")
		return ""
	}

	names, err := net.LookupAddr(ip.String())
	if err != nil {
		//fmt.Println("Error al buscar el dominio:", err)
		return ""
	}
	if len(names) > 0 {
		return names[0]
	}

	return ""
}
