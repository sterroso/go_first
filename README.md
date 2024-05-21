# A basic Frontend Masters Golang Course

This is my working project.

## Installing Go

1. Download the right Go version from the
   [download page](https://go.dev/dl/) and save it to your home directory.
   
   The file name might look something like: go<_version_>.<_platform_>.<_ext_>,
   where _version_ is the numeric version (e.g. `1.22.3`), _platform_ is
   the Operatoing System (OS) and processor type (e.g. `linux-amd64`) and
   _ext_ is the corresponding file extension (`zip`, `msi`, `tar.gz` or `pkg`)

   In this case, for simplicity's sake I'll assume you are installing Go
   version 1.22.3 on Ubuntu Linux 22.04 running over a 64-bit AMD/Intel PC,
   using a `tar.gz` installation file. So, my file of choice will be called
   `go1.22.3.linux-amd64.tar.gz`.

   - Option 1: using the `wget` command:
   ```
   $ wget -O go1.22.3.linux-amd64.tar.gz https://go.dev/dl/go1.22.3.linux-amd64.tar.gz
   ```

   - Option 2: using the `curl` command:
   ```
   $ curl -O https://go.dev/dl/go1.22.3.linux-amd64.tar.gz
   ```

   Optionally, you can verify the file's checksum by running the appropriate
   command:
   ```
   $ sha256sum go1.22.3.linux-amd64.tar.gz
   ```

2. Remove any previous version from the `/usr/local` and home directory:
   ```
   $ sudo rm -rf /usr/local/go && rm -rf ~/go
   ```
3. Unzip your recently donwnloaded file into the `/usr/local` directory:
   ```
   $ sudo tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz
   ```
4. Make sure to have the Go path in your `.profile`, `.bashrc` or `.zshrc`
   file. Add the following lines if not already present:
   ```
   # Set the Go PATH only if the directory exists
   if [ -d "/usr/local/go" ] ; then
       export PATH=$HOME/go/bin:/usr/local/go/bin:$PATH
   fi
   ```
5. Load the Go path by sourcing the corresponding file:
   ```
   $ source ~/.profile
   ```
6. Verify you Go installation:
   ```
   $ go version
   ```
7. Optionally, cleanup downloaded file
   ```
   $ rm go1.22.3.linux-amd64.tar.gz
   ```