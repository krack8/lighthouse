import { Component, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { Terminal } from 'xterm';
import { ActivatedRoute } from '@angular/router';
import { RequesterService } from '@core-ui/services/requester.service';
import { NgTerminal } from 'ng-terminal';

@Component({
  selector: 'kc-k8s-terminal',
  templateUrl: './k8s-terminal.component.html',
  styleUrls: ['./k8s-terminal.component.scss']
})
export class K8sTerminalComponent implements OnInit, OnDestroy {

  @ViewChild('term', { static: true }) child: NgTerminal;
  cols = parseInt(String(document.body.clientWidth / 9));
  rows = parseInt(String(document.body.clientHeight / 15.9));
  underlying: Terminal;
  connectContainer = false;
  appVpcId: string;
  pod: string;
  requester: any;
  domain: string;
  containerName;
  clusterId;
  namespace;
  socket: WebSocket;
  firstRequest = true;
  requestId: string;
  retryTimeout: any;

  constructor(private route: ActivatedRoute, private requesterService: RequesterService) {
    this.requester = this.requesterService.get();
    this.route.paramMap.subscribe(params => {
      // @ts-ignore
      //this.appVpcId = params.params.appVpcId;
      // @ts-ignore
      //this.pod = params.params.pod;
    });
  }

  ngOnInit() {
    this.route.queryParams.subscribe(res => {
      console.log(res.containerName, 'CONTAINER_NAME');
      this.containerName = res.containerName;
      this.domain = atob(res.domain);
      this.clusterId = res.clusterId;
      this.namespace = res.namespace;
      this.pod = res.pod;
      this.wsConnect();
    });
  }

  ngOnDestroy(): void {
    this.socket?.close();
    clearTimeout(this.retryTimeout);
  }

  wsConnect() {
    const that = this;
    const term = new Terminal({
      cursorBlink: true,
      rows: this.rows,
      cols: this.cols,
    });
    term.onData((data) => {
      if (this.socket.readyState === WebSocket.OPEN) {
        this.socket.send(data);
        term.focus();
      }
    });

    // Use WebSocket API
    let url = this.domain + '?token=' + this.requester.token + '&name=' + this.pod + '&container=' + this.containerName + '&namespace=' + this.namespace + '&cluster_id=' + this.clusterId + '&rows=' + this.rows + '&cols=' + this.cols;    console.log("ws URL  ----> ", url);   
    if (this.requestId) {
      url += '&taskId=' + this.requestId;
    }
    this.socket = new WebSocket(url);
    console.log('SOCKET', this.socket);

    const termDiv = document.getElementById('terminal-container');
    termDiv.innerHTML = '';
    term.write('\x1B[7;1;34m [Hello From KloverCloud!] \x1B[0m \n');
    term.open(document.getElementById('terminal-container'));

    this.socket.onmessage = (e) => {
      if (this.firstRequest) {
        this.firstRequest = false;
        this.requestId = e.data;
        console.log('Retry Request ID', this.requestId);
      } else {
        this.connectContainer = false;
        this.resize(this.socket, term);
        term.write(e.data);
      }
    };

    this.socket.onclose = (e) => {
      if (e.wasClean == true && e.code == 1000) {
        term.write('\x1B[3;1;31m Session Is Closed!\x1B[0m');
        this.connectContainer = true;
        this.firstRequest = true;
        this.requestId = null;
      } else {
        term.write('\x1B[3;1;31m Session Is Closed! Reconnecting...\x1B[0m');
        this.connectContainer = true;
        // Retry connection after a delay
        this.retryTimeout = setTimeout(() => {
          this.wsConnect();
        }, 5000); // Retry after 5 seconds
      }
    };

    this.socket.onopen = () => {
      that.resize(this.socket, term);
    };

    window.onresize = () => {
      // @ts-ignore
      that.resize(this.socket, term);
    };
  }

  resize(socket, term) {
    this.cols = parseInt(String(document.body.clientWidth / 9));
    this.rows = parseInt(String(document.body.clientHeight / 15.9));
    term.resize(this.cols, this.rows);
    if (socket.readyState === WebSocket.OPEN) {
      socket.send('{"cols":' + this.cols + ',"rows":' + this.rows + '}');
    }
  }
}