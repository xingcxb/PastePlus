# PastePlus

一个简单好用的📋粘贴板应用，由[不器](https://xingcxb.com?from=xingcxb/PastePlus)基于[wails](https://github.com/wailsapp/wails)开发。

## 简介

`PastePlus` 是一个简单好用的📋粘贴板应用，它可以帮助你快速的复制粘贴内容，同时还可以帮助你管理你的粘贴板历史。

## 功能

- [x] 粘贴板历史自动保存（仅文本）
- [ ] 粘贴板历史记录搜索
- [ ] 粘贴板历史记录删除
- [ ] 粘贴板历史记录清空
- [ ] 粘贴板历史记录导出
- [ ] 粘贴板配置

## 待处理问题

- [ ] 获取`pid`的时，如果没有激活过上一个应用会导致`panic`

## 截图

...

## 编译

- `macOS`

  `wails wails3 task package:darwin`
- `Windows`

  `wails wails3 task package:windows`
- `Linux`

  `wails wails3 task package:linux`

## 安装

### Windows

下载`exe`文件，双击运行即可。

### MacOS

下载`dmg`文件，双击运行即可。


## 使用

...

## 声明

本项目并不能承诺本项目绝对无漏洞，并且本项目也不会因漏洞造成的任何损失负责。

**如果不能理解或者接受，请不要使用本项目。**谢谢，选择是双向的，爱也是。

## 常见问题

- 它是什么？
  
    一个简单好用的📋粘贴板应用。
- 它有什么作用
  
    系统默认的剪切板仅仅存储的是您最后一个复制的内容，一旦您复制了新的内容，以前的的数据就会丢失。使用 `PastePlus` 复制的内容会自动保存到历史记录中，您可以随时查看以前复制的内容。
- 它是如何工作的？
  
    `PastePlus` 会在您复制内容的时候自动保存到历史记录中，您可以随时查看以前复制的内容。
- 如何查看历史记录？
  
    - 通用：点击 `PastePlus` 的图标选择第一个选项即可。
    - `Windows`: `Ctrl + Shift + V`
    - `Linux`: `Ctrl + Shift + V`
    - `MacOS`: `Command + Shift + V`

## 文档

[wails V3文档](https://v3alpha.wails.io/)

[MiSans](https://hyperos.mi.com/font)


## 致谢

[wails](https://github.com/wailsapp/wails/)

[Ant Design Vue](https://antdv.com/components/overview-cn)

[gohook](https://github.com/robotn/gohook)

[macdriver](https://github.com/progrium/macdriver)

[Icons8](https://igoutu.cn/)