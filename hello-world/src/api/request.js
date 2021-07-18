import api from './api'
import request from './axios'


const prefix = '/api/v1'

// 注册
export  function register(data){
    return new Promise((resolve, reject) => {
        request({
            url:  api+ '/register',
            method: 'post',
            data
        }).then((res) => {
            resolve(res)
        }).catch((err) => {
            reject(err)
        })
    })
} 
// 登录
export  function login(data){
    return new Promise((resolve, reject) => {
        request({
            url:  api + '/login',
            method: 'post',
            data
        }).then((res) => {
            resolve(res)
        }).catch((err) => {
            reject(err)
        })
    })
} 




//////////////////////////////////////

// 根据path查询游戏
export  function getGameByfullPath(params){
    return new Promise((resolve, reject) => {
        request({
            url:  api +prefix+ '/games/list',
            method: 'get',
            params
        }).then((res) => {
            resolve(res)
        }).catch((err) => {
            reject(err)
        })
    })
} 

// 根据id查询游戏
export  function getGameByID(params){
    return new Promise((resolve, reject) => {
        request({
            url:  api +prefix+ '/games/list/'+params.id,
            method: 'get',
            // params
        }).then((res) => {
            resolve(res)
        }).catch((err) => {
            reject(err)
        })
    })
} 
// 游戏列表
export  function listGames(data){
    return new Promise((resolve, reject) => {
        request({
            url:  api +prefix+ '/games/list',
            method: 'get',
            data
        }).then((res) => {
            resolve(res)
        }).catch((err) => {
            reject(err)
        })
    })
} 
// 用户添加游戏
export  function userAddGames(data){
    return new Promise((resolve, reject) => {
        request({
            url:  api +prefix+ '/usergame/add',
            method: 'post',
            data
        }).then((res) => {
            resolve(res)
        }).catch((err) => {
            reject(err)
        })
    })
} 
// 用户删除游戏
export  function userDeleteGames(data){
    return new Promise((resolve, reject) => {
        request({
            url:  api +prefix+ '/usergame/delete',
            method: 'post',
            data
        }).then((res) => {
            resolve(res)
        }).catch((err) => {
            reject(err)
        })
    })
} 
// 用户添加游戏列表
export  function userGamesList(params){
    return new Promise((resolve, reject) => {
        request({
            url:  api +prefix+ '/usergame/list',
            method: 'get',
            params
        }).then((res) => {
            resolve(res)
        }).catch((err) => {
            reject(err)
        })
    })
} 
// 用户排序游戏列表
export  function userOrderGames(data){
    return new Promise((resolve, reject) => {
        request({
            url:  api +prefix+ '/usergame/order',
            method: 'post',
            data
        }).then((res) => {
            resolve(res)
        }).catch((err) => {
            reject(err)
        })
    })
} 