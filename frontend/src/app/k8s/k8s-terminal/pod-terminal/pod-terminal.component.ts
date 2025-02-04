import { Component, OnDestroy, OnInit, ViewChild, ViewEncapsulation } from '@angular/core';
import { Terminal } from 'xterm';
import * as SockJS from 'sockjs-client';
import { NgTerminal } from 'ng-terminal';
import { ActivatedRoute } from '@angular/router';
import { RequesterService } from '@core-ui/services/requester.service';

@Component({
  selector: 'kc-pod-terminal',
  templateUrl: './pod-terminal.component.html',
  styleUrls: ['./pod-terminal.component.scss']
})
export class PodTerminalComponent implements OnInit, OnDestroy {
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
  socket: any;

  constructor(
    private route: ActivatedRoute,
    private requesterService: RequesterService
  ) {
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
  }

  wsConnect() {
    const that = this;
    const term = new Terminal({
      cursorBlink: true,
      rows: this.rows,
      cols: this.cols
    });
    term.onData(data => {
      if (this.socket.readyState === 1) {
        this.socket.send(data);
        term.focus();
      }
    });
    const url =
      this.domain +
      '/terminal/ws?' +
      'token=' +
      this.requester.token +
      '&pod=' +
      this.pod +
      '&containerName=' +
      this.containerName +
      '&namespace=' +
      this.namespace +
      '&clusterId=' +
      this.clusterId +
      '&rows=' +
      this.rows +
      '&cols=' +
      this.cols;
    console.log(this.domain, url);
    // Generate socket object
    this.socket = new SockJS(url);
    const termDiv = document.getElementById('terminal-container');
    termDiv.innerHTML = '';
    term.write('\x1B[7;1;34m [Hello From KloverCloud!] \x1B[0m \n');
    term.open(document.getElementById('terminal-container'));
    this.socket.onmessage = e => {
      this.connectContainer = false;
      this.resize(this.socket, term);
      term.write(e.data);
    };
    this.socket.onclose = e => {
      term.write('\x1B[3;1;31m Session Is Closed! \x1B[0m');
      this.connectContainer = true;
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
    if (socket.readyState === 1) {
      socket.send('{"cols":' + this.cols + ',"rows":' + this.rows + '}');
    }
  }
}
