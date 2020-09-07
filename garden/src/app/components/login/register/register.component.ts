import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.scss'],
})
export class RegisterComponent implements OnInit {
    constructor(public http: HttpClient) {}

    ngOnInit(): void {
        this.testGetdata();
    }
    private async testGetdata() {
        let res = await this.http.get(
            'http://47.99.205.217/huluxia/api/v1/articles/list?page=2&page_size=10'
        );
        console.log(res);
    }
}
