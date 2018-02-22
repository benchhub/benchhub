import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DashComponent } from './dash.component';
import {JobComponent} from "./job/job.component";
import {AboutComponent} from "./about/about.component";

const routes: Routes = [
  {
    path: 'dash',
    component: DashComponent,
    children: [
      {
        path: 'job',
        component: JobComponent
      },
      {
        path: 'about',
        component: AboutComponent
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class DashRoutingModule { }
