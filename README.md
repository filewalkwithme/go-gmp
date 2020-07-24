# go-openvas-gmp

Library to interact with Openvas 7.0 using the [gmp protocol](https://docs.greenbone.net/API/GMP/gmp-9.0.html) (Greenbone Management Protocol, `version 9.0`)

Here you willl find methods to create tasks, targets, scanners and retrieve results.

* This library is compatible with the latest version of Openvas: 7.0

# Documentation
- https://godoc.org/github.com/filewalkwithme/go-openvas-gmp/pkg/9/gmp
- https://godoc.org/github.com/filewalkwithme/go-openvas-gmp/pkg/9/gmp/client
- https://godoc.org/github.com/filewalkwithme/go-openvas-gmp/pkg/9/gmp/connections

# Sample Usage

```


// Connect to GVMD
conn, err := connections.NewUnixConnection("/tmp/openvas-socks/gvmd.sock")
if err != nil {
    panic(err)
}
defer conn.Close()

// Instantiate  a new GMP Client
gmpClient := client.New(conn)

// Authenticate
auth := &gmp.AuthenticateCommand{}
auth.Credentials.Username = "openvas"
auth.Credentials.Password = "openvas"
_, err = gmpClient.Authenticate(auth)
if err != nil {
    panic(err)
}

// Create a new task
newTask := &gmp.CreateTaskCommand{}
newTask.Name = "New Task"

newTask.Config = new(gmp.CreateTaskConfig)
newTask.Config.ID = "b9407b88-7b3c-47f3-a684-3605db80e5fd"

newTask.Target = new(gmp.CreateTaskTarget)
newTask.Target.ID = "a2a964f8-daa7-4e1c-ae3a-a72f06d49dbe"

newTask.Scanner = new(gmp.CreateTaskScanner)
newTask.Scanner.ID = "8abd321a-2eb1-4a7a-a368-fc118dc99a85"

newTaskResp, err := gmpClient.CreateTask(newTask)
if err != nil {
    panic(err)
}

// Start the task
st := &gmp.StartTaskCommand{}
st.TaskID = newTaskResp.ID
_, err = gmpClient.StartTask(st)
if err != nil {
    panic(err)
}
```

# Sample Application

First, you will need to start an Openvas instance. For this example we will use [Openvas-in-the-box](https://github.com/filewalkwithme/openvas-in-the-box.git), a ready to use Openvas Docker image. We are going to expose the GVMD Unix Socket under `/tmp/openvas-socks/gvmd.sock`

```
git clone https://github.com/filewalkwithme/openvas-in-the-box.git
sudo docker build openvas-in-the-box -t openvas
sudo docker run -d --rm -ti -p 80:80 -p 443:443 -v /tmp/openvas-socks:/var/run --name openvas openvas
```

Wait some seconds until `/tmp/openvas-socks/gvmd.sock` gets created by Openvas.
```
sleep 60
ls /tmp/openvas-socks/gvmd.sock
```

Next, we will execute the sample application avaible under the `examples` folder:
```
cd examples
go build -o go-openvas-gmp
./go-openvas-gmp
```

The sample application will generate an output like this:

```
Connecting to GVMD...
-> ok
Instantiate  a new GMP Client
-> ok
Authenticating...
-> ok
Getting the Default scanner...
-> ok (scanner id: 08b69003-5fc2-4037-a479-93b440211c73)
Getting the configuration named "Full and fast"...
-> ok (config id: daba56c8-73ec-11df-a475-002264764cea)
Creating a new target...
-> ok (target id: 34e34484-2dfa-4f01-b9fe-2b90e4a38fba)
Creating a new task...
-> ok (task id: d82a368f-acfa-422e-a78a-63727861f5a9)
Start the task

Monitoring task progress...
Monitoring task progress: 1%
Monitoring task progress: 2%
Monitoring task progress: 6%
Monitoring task progress: 8%
Monitoring task progress: 14%
Monitoring task progress: 52%
Monitoring task progress: 80%
Monitoring task progress: 90%
Monitoring task progress: 98%
Monitoring task progress: 100%


Result[0]: CGI Scanning Consolidation (score: 0.0)
...
Result[7]: HTTP Security Headers Detection (score: 0.0)
...
Result[15]: Services (score: 0.0)
...
Result[20]: SSL/TLS: Diffie-Hellman Key Exchange Insufficient DH Group Strength Vulnerability (score: 4.0)
Result[37]: SSL/TLS: Report Vulnerable Cipher Suites for HTTPS (score: 5.0)
Result[38]: SSL/TLS: Report Weak Cipher Suites (score: 4.3)
Result[39]: Traceroute (score: 0.0)
Result[40]: Unknown OS and Service Banner Reporting (score: 0.0)
```

# Tests
```
go test ./... -coverprofile=coverage.out

?       github.com/filewalkwithme/go-openvas-gmp/examples       [no test files]
ok      github.com/filewalkwithme/go-openvas-gmp/pkg/9/gmp      0.004s  coverage: [no statements]
ok      github.com/filewalkwithme/go-openvas-gmp/pkg/9/gmp/client       0.017s  coverage: 100.0% of statements
ok      github.com/filewalkwithme/go-openvas-gmp/pkg/9/gmp/connections  0.009s  coverage: 100.0% of statements
ok      github.com/filewalkwithme/go-openvas-gmp/pkg/9/gmp/connections/internal/implementation  0.003s  coverage: 100.0% of statements
```

HTML Report:
```
go tool cover -html=coverage.out
```
