import { Component, OnInit } from '@angular/core';
import { NodeService } from './node.service';

@Component({
  selector: 'app-node',
  templateUrl: './node.component.html',
  styleUrls: ['./node.component.css']
})
export class NodeComponent implements OnInit {

  constructor(private svc: NodeService) {
  }

  ngOnInit() {
    // TODO: it would be good if we can generate ts models ... based on protobuf ...
    this.svc.getCentral()
      .subscribe(data => {
        console.log(data)
      }, err => {
        console.error(err)
      })
  }

}
