import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {DashComponent} from './dash.component';
import {JobComponent} from "./job/job.component";
import {AboutComponent} from "./about/about.component";
import {NodeComponent} from "./node/node.component";

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
        path: 'node',
        component: NodeComponent
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
export class DashRoutingModule {
}
