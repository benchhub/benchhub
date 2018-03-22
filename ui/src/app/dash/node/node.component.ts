import { Component, OnInit } from '@angular/core';
import { NodeService } from './node.service';

@Component({
  selector: 'app-node',
  templateUrl: './node.component.html',
  styleUrls: ['./node.component.css']
})
export class NodeComponent implements OnInit {
  central;
  agents;
  agentsCount = 0;

  constructor(private svc: NodeService) {
  }

  ngOnInit() {
    // TODO: it would be good if we can generate ts models ... based on protobuf ...
    this.svc.getCentral()
      .subscribe(data => {
        console.log('got central node', data);
        // TODO: need model for mapping the data
        // this.central = data.node;
      }, err => {
        console.error(err)
      });
    this.svc.getAgents()
      .subscribe(data => {
        console.log('got agents nodes', data);
        // this.agents = data.agents;
        // this.agentsCount = data.agents.length;
      }, err => {
        console.error(err)
      })
  }

}
