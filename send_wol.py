import socket
import sys
import binascii

IP = "255.255.255.255"
PORT = 7

def convert_mac(mac: str) -> bytearray:
    mac = mac.replace(":", "")
    return bytearray.fromhex(mac)

def prepare_packet(mac: bytearray) -> bytearray:
    ar = bytearray(6 * 16 + 6)
    ar[0] = 0xFF
    ar[1] = 0xFF
    ar[2] = 0xFF
    ar[3] = 0xFF
    ar[4] = 0xFF
    ar[5] = 0xFF

    for i in range(0, 6 * 16):
        ar[i + 6] = mac[i % 6]

    return ar

if __name__ == '__main__':
    if len(sys.argv) < 2:
        print("must specify MAC address. (ex. \"B0:83:FE:6A:A7:C9\")")
        exit(1)
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM) # Ca serait cool d'utiliser un paquet plus bas lvl que UDP si possible
    sock.setsockopt(socket.SOL_SOCKET, socket.SO_BROADCAST, 1)
    mac = convert_mac(sys.argv[1])
    packet = prepare_packet(mac)
    print(binascii.hexlify(packet))
    sock.sendto(packet, (IP, 7))