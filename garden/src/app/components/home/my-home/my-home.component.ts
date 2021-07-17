import { Component, OnInit } from '@angular/core';

import { HttprequestService } from '../../../services/http-request.service';

@Component({
    selector: 'app-my-home',
    templateUrl: './my-home.component.html',
    styleUrls: ['./my-home.component.scss'],
})
export class MyHomeComponent implements OnInit {
    constructor(private http: HttprequestService) {}
    userinfo: any = {};
    ngOnInit(): void {
        this.userinfo = JSON.parse(localStorage.getItem('loginUser'));
    }
}
