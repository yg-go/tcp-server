
import socket
import struct
import datetime


def makeHeader(datasize, opcode):
    return struct.pack("hh", datasize, opcode)

def makeTrackInfo(numTrack):
    l = []
    for i in range(numTrack):
        trackid = i
        x = i
        y = i
        z = i
        power = i
        l.append(struct.pack("iffff", trackid, x, y, z, power))

    ret = struct.pack("i", numTrack)
    for i in l:
        ret += i
    return ret


def makePlotInfo(numPlot):
    l = []
    for i in range(numPlot):
        plotid = i
        x = i
        y = i
        z = i
        power = i
        l.append(struct.pack("iffff", plotid, x, y, z, power))

    ret = struct.pack("i", numPlot)
    for i in l:
        ret += i
    return ret

#HURA1 PROTOCOL TEST
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.connect(("192.168.0.191", 10050))

total = b""
#Track Test ============================================================================================================
tracks = makeTrackInfo(2)
header = makeHeader(len(tracks)+4, 0x07D6)
payload = header+tracks
total += payload

plots = makePlotInfo(2)
header = makeHeader(len(plots)+4, 0x07D5)
payload = header+plots
total += payload

sock.sendall(total)
sock.close()
