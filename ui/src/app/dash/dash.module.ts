import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { NgZorroAntdModule } from 'ng-zorro-antd';

import { DashRoutingModule } from './dash-routing.module';
import { DashComponent } from './dash.component';

@NgModule({
  imports: [
    CommonModule,
    DashRoutingModule,
    NgZorroAntdModule
  ],
  declarations: [DashComponent]
})
export class DashModule { }
