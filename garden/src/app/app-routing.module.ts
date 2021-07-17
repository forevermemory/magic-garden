import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// 自定义的组件
import { RegisterComponent } from './components/login/register/register.component';
import { LoginComponent } from './components/login/login/login.component';
import { MyHomeComponent } from './components/home/my-home/my-home.component';

// 配置路由组
const routes: Routes = [
    { path: 'login/register', component: RegisterComponent },
    { path: 'login/login', component: LoginComponent },
    { path: 'home/my_home', component: MyHomeComponent },
    { path: '**', redirectTo: 'login/register' }, // 重定向 默认路由
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule],
})
export class AppRoutingModule {}
