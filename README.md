# oneindex-typora-uploader

## 介绍

+ 一个自定义的typora图片上传程序
+ 需自备go环境

## 使用

1. 克隆整个仓库

   ```
   git clone https://github.com/duskLight-wang/oneindex-typora-uploader.git
   ```

2. 用vscode打开项目，更改`main.go`中的`tar_url`

   ![image-20211008120535616](http://cloud.dusklight.top/images/2021/10/08/K5RVZTBjm9/image-20211008120535616.png)
3. 编译项目
   
   ```
   go build
   ```
4. 配置oneindex

   ![image-20211008121015163](http://cloud.dusklight.top/images/2021/10/08/RnuVfdkTqU/image-20211008121015163.png)

5. 配置typora
 
   + 将刚才编译生成的exe文件的存放路径放到自定义命令的路径中
   ![image-20211008121357041](http://cloud.dusklight.top/images/2021/10/08/xTIJ7W6el1/image-20211008121357041.png)

6. 直接粘贴图片或拖动到typora编辑区域，图片将自动完成上传并且引用

## Tips

+ typora存在测试bug，无法同时上传多图片，单图片上传正常(非本程序问题)
+ 建议配合[snipaste](https://www.snipaste.com/)使用
+ 使用世纪互联，图片加载会更加快速
+ onedrive有翻车风险，建议不会再更改的文件导出为pdf

