import sys
import paramiko

if __name__ == '__main__':
    if len(sys.argv) < 2:
        print("must specify username & address. (ex. nuagemalin@20.126.66.3)")
        exit(84)
    args = sys.argv[1].split('@')
    client = paramiko.SSHClient()
    client.load_system_host_keys()
    client.connect(args[1], username=args[0])
    stdin, stdout, stderr = client.exec_command("cat /sys/power/state")
    stdin.close()
    print(stdout.read())