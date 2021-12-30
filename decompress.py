import zlib


def main():
    all()
    stream()


def all():
    print("==== RUNNING DECOMPRESSION OF all.in")

    d_obj = zlib.decompressobj()

    with open("chunks/all.in", "rb") as f:
        data = f.read()

    out = d_obj.decompress(data)
    print(len(out))
    print(out)
    print("")


def stream():
    print("==== RUNNING DECOMPRESSION OF SEPARATE CHUNKS")
    print("")

    d_obj = zlib.decompressobj()

    with open("chunks/0.in", "rb") as f:
        data = f.read()

    out = d_obj.decompress(data)
    print(len(out))
    print(out)

    with open("chunks/1.in", "rb") as f:
        data = f.read()

    out = d_obj.decompress(data)
    print(len(out))
    print(out)


if __name__ == "__main__":
    main()
