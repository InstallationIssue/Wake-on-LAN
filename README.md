# Wake-on-LAN

This is a quick go project, designed to create an executable to wake systems on LANs using magic packets.

## To Build and Run

### Before Building

This was built and tested using Mac and Linux, will still work if used to wake windows systems but if building the executable on windows, some commands in the following steps may ahve to be changed.

Additionally, not all network cards support wake on lan (particularly wireless cards) and magic packets cannot be used to wake systems if not on the same LAN.

To check if your network card supports WOL, you may need to check your systems BIOS.

### Follow These Steps

1. Get the MAC address of the network card from the system you want to wake over LAN.

   On Macs and Linux, you can enter the following on the command line and search for your network interface:

   ```
   ifconfig
   ```

   Alternatively, you can inspect your network details through system preferences on Macs or System settings on graphical Linux distros

   The MAC address should have a form similar to below

   **_mac = A1:2B:C3:4D:E5:6F_**

2. Get the broadcast address of your LAN network

   This will be the last IP address in your IP space.

   This depends on your local network configuration, but for a home network with a typical IP address of 192.168.1.0 and subnet mask of 255.255.255.0, the broadcast address would be 192.168.1.255

   **_broadcast-ip = 192.168.1.255_**

3. Choose a udp broadcast port
   This can be any available port on the router this script will be sending the signal to.

   This does not usually matter but to be safe if some ports are reserved, typically most ports in the 8000+ range are free

   **_port = 8001_**

4. Install Go 1.23.0 if not installed. Later versions may work but this was developerd using v1.23.0

5. Build the executable using the following command, replacing all fields between <> with your own values

   ```
   go build -o <name-of-executable> -ldflags="-X main.macTarget=<mac> -X main.ipBroadcast=<broadcast-ip> -X main.port=<port>" <path-to-this-directory>
   ```

   For example...

   ```
   go build -o wol -ldflags="-X main.macTarget=A1:2B:C3:4D:E5:6F -X main.ipBroadcast=192.168.1.255 -X main.port=8001" .
   ```

6. Execute the built executable
