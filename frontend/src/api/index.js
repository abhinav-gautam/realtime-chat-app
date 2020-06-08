var socket = new WebSocket("ws://localhost:8080/ws")

const connect = cb =>{
    console.log("[+] Attempting Connection...")
    socket.onopen =()=>{
        console.log("[+] Successfully Connected")
    }
    socket.onopen =msg=>{
        console.log("[+] Connected")
        console.log(msg)
    }
    socket.onerror=error=>{
        console.log("[-] Socket Error -->"+error)
    }
    socket.onclose=event=>{
        console.log("[-] Socket Connection Closed -->"+event)
    }
    socket.onmessage = msg =>{
        console.log("[+] Message Received")
        console.log(msg)
        cb(msg)
    }
}
const sendMsg =msg=>{
    console.log("[+] Sending Message :"+msg)
    socket.send(msg)
}
export {connect,sendMsg}