package main

import (
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
	// "github.com/mdp/qrterminal/v3"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
)

func main() {
	/* get computer info */
	var serialNumber string
	var model string
	var firmwareVersion string

	// ioreg requires macOS 10.12 or newer
	// system_profiler works since 10.6 at least
	// reference: https://stackoverflow.com/a/44718088
	out, err := exec.Command("/usr/sbin/system_profiler", "SPHardwareDataType").Output()
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range strings.Split(string(out), "\n") {

		// reference: https://osxdaily.com/2011/04/25/get-mac-serial-number-command-line/
		if strings.Contains(line, "Serial Number (system)") {
			s := strings.Split(line, " ")
			serialNumber = s[len(s)-1]
		}

		// reference: https://www.tardisk.com/pages/how-to-check-your-firmware-version
		if strings.Contains(line, "Model Identifier") {
			s := strings.Split(line, " ")
			model = s[len(s)-1]
		}

		// reference: https://www.tardisk.com/pages/how-to-check-your-firmware-version
		if strings.Contains(line, "System Firmware Version") || strings.Contains(line, "Boot ROM Version") {
			s := strings.Split(line, " ")
			firmwareVersion = s[len(s)-1]
		}
	}
	log.Infof("Detected Serial Number: %s", serialNumber)
	log.Infof("Detected Model: %s", model)
	log.Infof("Detected Firmware Version: %s", firmwareVersion)

	// create accessory
	info := accessory.Info{
		Name:             "Computer",
		SerialNumber:     serialNumber,
		Manufacturer:     "Apple",
		Model:            model,
		FirmwareRevision: firmwareVersion,
	}
	tv := accessory.NewTelevision(info)

	// add characteristics
	tv.Television.Active.SetValue(characteristic.ActiveActive)
	tv.Television.Active.OnValueRemoteUpdate(func(v int) {
		log.Infof("active => %d\n", v)
	})

	// config ip transport
	config := hc.Config{Pin: "11346789", StoragePath: "./hkd_db"}
	t, err := hc.NewIPTransport(config, tv.Accessory)
	if err != nil {
		log.Fatal(err)
	}
	// uri, _ := t.XHMURI()
	// qrterminal.Generate(uri, qrterminal.L, os.Stdout)

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}
