import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { FormsModule } from '@angular/forms'; // 双向数据绑定

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
// 服务
import { HttprequestService } from './services/http-request.service';

// 组件库
import { BreadcrumbModule } from 'primeng/breadcrumb';
// import { AccordionModule } from 'primeng/accordion'; //accordion and accordion tab

// 自定义组件
import { HeaderComponent } from './components/header/header.component';
import { FooterComponent } from './components/footer/footer.component';
import { RegisterComponent } from './components/login/register/register.component';
import { LoginComponent } from './components/login/login/login.component';
import { HttpClientModule } from '@angular/common/http';

@NgModule({
    declarations: [
        AppComponent,
        HeaderComponent,
        FooterComponent,
        RegisterComponent,
        LoginComponent,
    ],
    imports: [
        BrowserModule,
        AppRoutingModule,
        HttpClientModule,
        FormsModule,
        // 组件库
        // AccordionModule,
        BreadcrumbModule,
    ],
    providers: [HttprequestService],
    bootstrap: [AppComponent],
})
export class AppModule {}
