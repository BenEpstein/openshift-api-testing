# Download the latest Go version
wget https://go.dev/dl/go1.23.1.linux-amd64.tar.gz

# Extract and install
sudo tar -C /usr/local -xzf go1.23.1.linux-amd64.tar.gz

# Set up environment variables in ~/.bashrc or ~/.profile
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc

# Apply the changes
source ~/.bashrc

# Verify the installation
go version
