Page({
    // 首页数据结构
    data: {
        // 滚动图片
        banner: [],
        channel: [],
        content: [],
        hidden: false
    },
    // 刚启动时加载
    onLoad: function (options) {
        var that = this;
        // 获取主页所需要的数据
        getIndexData(that);
    },
    // 点击每个HotMsg时触发
    onItemClick: function (event) {
        var targetUrl = Constant.PAGE_SPECIFIC;
        if (event.currentTarget.dataset.publishTime != null)
            targetUrl = targetUrl + "?publishTime=" + event.currentTarget.dataset.publishTime;
        wx.navigateTo({
            url: targetUrl
        });
    },
    // 点击提交数据时，触发
    onPostClick: function (event) {
        console.log("onPostClick");
        wx.navigateTo({
            url: Constant.PAGE_POSt
        });
    }
});

/**
 * 获取主页所需要的数据
 * @param that Page的对象，用其进行数据的更新
 */
function getIndexData(that) {
    wx.request({
        url: Constant.INDEX_URL,
        header: {
            "Content-Type": "application/json"
        },
        success: function (res) {
            if (res == null || res.data.code != 0) {
                console.error("get index data failed");
                wx.showModal({
                    title: 'Warning',
                    content: '加载失败，请重试~',
                    success: function (r) {
                        console.log("try again");
                    }
                });
                return;
            }
            console.log(res.data.data);
            that.setData({
                banner: res.data.data.banner,
                channel: [
                    {
                        "target": "/pages/main/main",
                        "icon_url": "https://api.lecury.cn/static/img/index/icon_1.png",
                        "name": "andriod",
                    },
                    {
                        "target": "/pages/main/main",
                        "icon_url": "https://api.lecury.cn/static/img/index/icon_1.png",
                        "name": "andriod",
                    },
                    {
                        "target": "/pages/main/main",
                        "icon_url": "https://api.lecury.cn/static/img/index/icon_1.png",
                        "name": "andriod",
                    },
                ],
                content: res.data.data.content,
                hidden: true
            });
        }
    });
}

var Constant = require('../../utils/constant.js');
