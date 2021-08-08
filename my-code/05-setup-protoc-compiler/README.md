# Installation
In order to perform code generation, you will need to install protoc  on your computer.

## MacOSX

It is actually very easy, open a command line interface and type 
```shell
brew install protobuf
```

## Ubuntu (Linux)

Find the correct protocol buffers version based on your Linux Distro: https://github.com/google/protobuf/releases

### Example with x64:
```shell
# Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
# Unzip
unzip protoc-3.5.1-linux-x86_64.zip -d protoc3
# Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/
# Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/
# Optional: change owner
sudo chown [user] /usr/local/bin/protoc
sudo chown -R [user] /usr/local/include/google
```

## Windows

1. Download the windows archive: https://github.com/google/protobuf/releases
   - Example: https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-win32.zip

2. Extract all to `C:\proto3`

    Your directory structure should now be
```shell
C:\proto3\bin
C:\proto3\include
```

3. Finally, add C:\proto3\bin to your PATH:

   - From the desktop, right-click **_My Computer_** and click **_Properties_**.
   - In the **_System Properties_** window, click on the **_Advanced_** tab
   - In the **_Advanced_** section, click the **_Environment Variables_** button.
   - Finally, in the _Environment Variables_ window (as shown below), highlight the Path variable in the Systems Variable section and click the **_Edit_** button. 
   - Add or modify the path lines with the paths you wish the computer to access. Each different directory is separated with a semicolon as shown below.
   
    `C:\Program Files; C:\Winnt; ...... ; C:\proto3\bin` (you need to add ; C:\proto3\bin  at the end)

---

# Code Generation

### generate Python code
```shell
protoc --proto_path=proto-files --python_out=python-out proto-files/city.proto
```
### generate Go code

- to be able to generate Go, we need to add `option go_package` in our proto file `city.proto`
```protobuf
option go_package = ".;gen";
```
- now generate, better we use absolute path:
```shell
protoc -I=$PWD/proto-files --go_out=$PWD/golang-out $PWD/proto-files/*.proto
```
