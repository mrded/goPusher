# goPusher
(SSE) Server-Sent Event Service

Usage example:

        curl -HContent-Type:application/json -d '{"name":"my_event","channel":"my_channel","data":"hello world","token":"secret"}' http://localhost:9090/events
