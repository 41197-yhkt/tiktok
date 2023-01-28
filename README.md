# tiktok
大家好，这里是艺画开天队的Tiktok大项目
这里是项目的总目录（根目录）
通过git submodule的方式，这里所有的子模块（详情可以参考git module官方文档: https://git-scm.com/book/zh/v2/Git-%E5%B7%A5%E5%85%B7-%E5%AD%90%E6%A8%A1%E5%9D%97）


# 研发方式
1. 根目录引用了所有依赖的子目录，所以每个微服务团队的programmer，cd到对应的目录之后，本质上就进入了自己负责的仓库，直接进行研发就好（可以不动别的仓库）。
2. 这个项目的目录，是能够监听所有子项目的git变化的。例如你现在如果正处于tiktok目录，如果你执行`cd idl`，那么本质上，你就是切换到了idl这个项目中，进行研发，和你本地git clone idl分支并研发的效果是一模一样的。研发完成，最终完成push即可。 
3. 处于主项目的programmer，通过git status，你可以看到各个子模块的情况（例如idl目录有新的提交）。如果你想更新某个目录（例如，你们开发小组，切了一个idl分支，并对于idl进行了改动和提交），你只需要cd到idl子模块中，完成分支切换和更新即可。
4. 推荐在主项目中，进入子项目模块进行研发，因为这种情况下，你可以使用别的子模块中的代码。（例如你在开发gateway模块，你的父目录是tiktok目录，你可以通过../idl，访问到idl模块的代码，从而很方便的，完成跨子模块代码的调用。
5. 如果你想要调用别的子模块逻辑，可能还需要进行一些依赖管理（孩子只用过gradle和maven，对于golang没有什么发言权。
6. 一些公用模块，如果想要复用，例如model，可以通过类似的方式，通过git submodule add的方法，添加子模块(可以自行Google)

# 第一次Clone tiktok项目时，您需要做什么？
克隆项目后，默认子模块目录下无任何内容。需要在项目根目录执行如下命令完成子模块的下载：
git submodule update --init --recursive
git submodule foreach -q --recursive 'git checkout $(git config -f $toplevel/.gitmodules submodule.$name.branch || echo main)'
> 参考来自于： https://zhuanlan.zhihu.com/p/421381523