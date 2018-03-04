import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { NgZorroAntdModule } from 'ng-zorro-antd';

import { DashRoutingModule } from './dash-routing.module';
import { DashComponent } from './dash.component';
import { JobComponent } from './job/job.component';
import { AboutComponent } from './about/about.component';
import { NodeComponent } from './node/node.component';
import { NodeService } from "./node/node.service";

@NgModule({
  imports: [
    CommonModule,
    DashRoutingModule,
    NgZorroAntdModule
  ],
  declarations: [DashComponent, JobComponent, AboutComponent, NodeComponent],
  providers: [NodeService]
})
export class DashModule {
}
