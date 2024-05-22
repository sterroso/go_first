# A basic _Frontend Masters_ Go Course

This is my working project.

## ⚙️ Installing Go

1. ⬇️ Download the right Go version from the
   [download page](https://go.dev/dl/) and save it to your home directory.
   
   The file name might look something like: go<_version_>.<_osarch_>.<_ext_>,
   where _version_ is the numeric version (e.g. _1.22.3_), _osarch_ is
   the Operating System (OS) and CPU architecture (e.g. _linux-amd64_ for a
   Linux OS over an 64-bit AMD/Intel architecture CPU) and _ext_ is the
   corresponding file extension (_zip_, _msi_, _tar.gz_ or _pkg_).

   In this case, for simplicity's sake I'll assume you are installing Go
   version _1.22.3_ on Ubuntu _Linux_ 22.04 OS, running over a _64-bit
   AMD/Intel PC_, using a _tar.gz_ installation file. So, my file of choice
   will be called `go1.22.3.linux-amd64.tar.gz`.

   As a Linux user, I can choose one of the next methods to download the file:

   - Option 1: using the **wget** command:
   ```
   wget -O go1.22.3.linux-amd64.tar.gz https://go.dev/dl/go1.22.3.linux-amd64.tar.gz
   ```

   - Option 2: using the **curl** command:
   ```
   curl -O https://go.dev/dl/go1.22.3.linux-amd64.tar.gz
   ```

   - Option 3: Use my web browser to navigate to the [official Go download
     page](https://go.dev/dl/) and click on the appropriate link.

   Optionally, you can verify the file's checksum by running the appropriate
   command after downloading the Go installation file:
   ```
   if [[ $(sha256sum go1.22.3.linux-amd64.tar.gz) = '8920ea521bad8f6b7bc377b4824982e011c19af27df88a815e3586ea895f1b36  go1.22.3.linux-amd64.tar.gz' ]] { print Ok } else { print 'Checksum does not match' }
   ```

2. ❌ Remove any previous version from the `/usr/local` and home directory:
   ```
   sudo rm -rf /usr/local/go && rm -rf ~/go
   ```
3. 📂 Unzip your recently donwnloaded file into the `/usr/local` directory:
   ```
   sudo tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz
   ```
4. 📝 Make sure to have the Go path in your `.profile`, `.bashrc` or `.zshrc`
   file. Add the following lines if not already present:
   ```
   # Set the Go PATH only if the directory exists
   if [ -d "/usr/local/go" ] ; then
       export PATH=$HOME/go/bin:/usr/local/go/bin:$PATH
   fi
   ```
5. 💻 Load the Go path by sourcing the corresponding file:
   ```
   source ~/.profile
   ```
6. ✅ Verify you Go installation:
   ```
   go version
   ```
7. 🧹 Optionally, cleanup downloaded file
   ```
   rm go1.22.3.linux-amd64.tar.gz
   ```