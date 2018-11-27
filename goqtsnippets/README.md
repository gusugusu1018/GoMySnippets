<img src="https://github.com/minaminoki/GoQtSnippets/blob/master/GopherChan.png" width="640">

# Overview
This repository is just my golang snippets for cross platform GUI application using Qt widgets.  
Qt godoc is [here](https://godoc.org/github.com/therecipe/qt).  

# Examples 
## PushButton
![img](https://github.com/minaminoki/GoQtSnippets/blob/master/PushButton/screenShot.png)

## Slider
![img](https://github.com/minaminoki/GoQtSnippets/blob/master/Slider/screenShot.png)

## GroupBox
![img](https://github.com/minaminoki/GoQtSnippets/blob/master/GroupBox/screenShot.png)

## ProgressBar
![img](https://github.com/minaminoki/GoQtSnippets/blob/master/ProgressBar/screenShot.png)

## Apply Style
Style sheet example documents is [here](http://doc.qt.io/qt-5/stylesheet-examples.html).  
Style sheet reference documents is [here](http://doc.qt.io/qt-5/stylesheet-reference.html).  
![img](https://github.com/minaminoki/GoQtSnippets/blob/master/ApplyStyle/screenShot.png)

## Chart
QtChart documents is [here](https://doc.qt.io/qt-5.11/qtcharts-index.html).
QtChart godoc is [here](https://godoc.org/github.com/therecipe/qt/charts).  

### Line Chart
![img](https://github.com/minaminoki/GoQtSnippets/blob/master/Charts/LineChart/screenShot.png)

### Dynamic Line Chart
![gif](https://raw.githubusercontent.com/wiki/minaminoki/GoQtSnippets/images/DynamicLineChart.gif)

## File Dialog  
![img](https://github.com/minaminoki/GoQtSnippets/blob/master/FileDialog/screenShot.png)  

## Image View  
Reference is [here](https://socketloop.com/tutorials/golang-qt-image-viewer-example).  
![img](https://github.com/minaminoki/GoQtSnippets/blob/master/ImageView/screenShot.png)  

## Splash Screen  
Reference is [here](https://socketloop.com/tutorials/golang-qt-splash-screen-with-delay-example).  

## Progress Dialog  
Reference is [here](https://socketloop.com/tutorials/golang-qt-progress-dialog-example).  
Need this line for update UI  
```
core.QCoreApplication_ProcessEvents(core.QEventLoop__AllEvents)
```

# How to run  
## Installation  
Read [here](https://github.com/therecipe/qt/wiki/Installation).  

## Run Example  
```
git clone https://github.com/minaminoki/GoQtSnippets/
cd GoQtSnippets/PushButton
```
and make or 
```
qtdeploy test desktop .
```
