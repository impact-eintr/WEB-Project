const num = 1;
let price = 1.5;
let myName = "yixingwei";
let trueValue = true;
let nullVal = null;
let und;
for (let i =0;i < 10;i++){
    document.write(i);
    document.write("<br>");
}
document.writeln("num= "+num);
document.write("<br>");
document.writeln(typeof num);
function Hello(stu){
    let hi = "hi";
    console.log(hi);
    console.log(stu)
}

//类的构造函数
function Stu(ID,name){
    this.ID = ID;
    this.name  = name;
    this.PrivateFunc = function(){
        alert("这是一个类内函数，可以可以声明为私有的");
    };
    
}

Stu.prototype.publicFunc = function(){
    document.write("这是一个类外函数，只能是公有的");
}

let s = new Stu(
    "2018006578",
    "yixingwei",
);


Hello(s);
s.publicFunc();

//let 与 var 的变量作用范围
let movie = "hhh";

function A(){
    const movie = "a";
    return movie;
}

function B(){
    const movie = "b";
    return movie;
}

//箭头函数
let testFunc = function(){
    return s.ID;
}

let ArrowFunc1 = () => s.ID;

let ArrowFunc2 = name => {
    s.name = name;
    return s.name;
}

ArrowFunc2("wanghao");
document.writeln(s.name);

//模板字面量
document.writeln(
    `tell me your name
    <br>
    ok,I'm ${s.name}`);

function sum(x = 1,y = 2,z = 3){
    return x+y+z;
}

//数组解构
let numlsit = [4,5,6];
console.log(sum(...numlsit));

//数组解构初始化多个变量
//let [x,y,z] = ['a','b',3]
//console.log(x,y,z)
//数组结构交换元素
let x = 9;
let y = 8;
[x,y] = [y,x]
document.write("<br>");
document.write(x,y);

//简写方法名
const hello = {
    name : "yixingwei",
    PrintHello(){
        document.write("<br>hello");
    }
};

hello.PrintHello();

//面向对象
class Book {
    constructor(title,pages,isbn){
        this.title = title;
        this.pages = pages;
        this.isbn = isbn;
    }
    printIsbn(){
        document.write("<br>",this.isbn);
    }
}

var book = new Book("A",33,676676);
book.printIsbn()
