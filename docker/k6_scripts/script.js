import http from 'k6/http'

export let options = {
    stages:[
        {target: 10, duration: '5s'},
        {target: 15, duration: '1m'},
        {target: 3, duration: '5s'},
    ]
}

export default function() {
    http.get("http://host.docker.internal:8009/api/devices")
}