import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

// const centralUrl = 'http://localhost:6091/api';
const centralUrl = '/api';

@Injectable()
export class NodeService {

  constructor(private http: HttpClient) {

  }

  getCentral() {
    return this.http.get(centralUrl + '/node')
  }

  getAgents() {
    return this.http.get(centralUrl + '/agent/list')
  }
}
