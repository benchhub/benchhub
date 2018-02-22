import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { NgZorroAntdModule } from 'ng-zorro-antd';

import { DashRoutingModule } from './dash-routing.module';
import { DashComponent } from './dash.component';
import { JobComponent } from './job/job.component';
import { AboutComponent } from './about/about.component';

@NgModule({
  imports: [
    CommonModule,
    DashRoutingModule,
    NgZorroAntdModule
  ],
  declarations: [DashComponent, JobComponent, AboutComponent]
})
export class DashModule { }
