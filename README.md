# Nephele

我们为Nephele文档提供的内容定义了4个需求等级(Requirement Class)。各个需求等级没有强依赖关系，但有道是：高级需求的产生是低级需求相对满足的结果。

**Class 1** 
「了解」需要精简的陈述以填补基本认知的空白。
标记为C1的内容通常是“是什么”，“能做什么”，“能做到什么程度”等常见问题的答案。

**Class 2** 
「使用」需要具体的要求与建议指导实战。
标记为C2的内容会是严肃的，全面的，细致的。这部分内容会占据整个文档群的绝大部分篇幅，有时会不可避免的高频的出现一些重复的文字。因此也会设置目录方便查阅。（还是建议有至少一次的完整浏览）

**Class 3* 
「理解」需要学习内在机制以满足好奇心，或借鉴并构建类似的，相关的或其他的系统。
标记为C3的内容通常会更加枯燥，冗长。它们可能会给读者一些启发与灵感，但也可能不会。对于那些把关注点放在实际应用的读者而言，这部分内容或许能够产生一些安全感吧，但总的来说，会显得不是那么重要。

**Class 4* 
「交流」需要参与贡献社区或者需要挑战作者的理念。
互相理解是合作的坚盾，知己知己是争论的长矛。标记为C4的内容是作者团队的心里话，会是他们的一些展望，也可能会是总结。这些内容一定是天马行空，充满活力的。


## Background

随着图片量的快速增长，需要专门的图片系统来管理...
随着业务的不断发展，图片预先处理已不能满足业务需求，需要访问时实时动态缩放裁剪等处理...
随着对图片质量要求越来越高，各种滤镜等高级需求增长，添加了图片编辑器，供用户自行编辑。

Nephele因此而诞生。

## Introduction

Nephele是携程框架部门研发的开源图片服务，使用Go开发，意在为业务和开发提供一套完整的在线图片解决方案，其架构简单，可以稳定运行在绝大多数主流环境下，不仅提供功能完善的图片上传，访问，处理api，可视化监控和在线图片编辑功能也能够方便非开发人员协同管理。

## Functionality
  * **多存储支持**
  	* 本地文件，FastDFS，OSS，S3
  * **图片编码及优化**
  	* JPEG，PNG，GIF，WEBP。
  	* 动态最优质量评估。
  * **多种图片处理**
  	* 裁剪，缩放，旋转，转换，滤镜，水印，原图保护，样式。
  * **图片监控**
  	* 资源消耗，访问量，延时，失败率，回源率。
  * **图片管理面板**
  	* 图片管理，图片编辑。  
  * **快速部署** 		
  	  	
  	
  	
  	
  	
