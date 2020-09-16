(window.webpackJsonp = window.webpackJsonp || []).push([
    [0],
    {
        1: function(t, e, n) {
            "use strict";
            n.d(e, "k", (function() {
                return y
            })), n.d(e, "m", (function() {
                return w
            })), n.d(e, "l", (function() {
                return x
            })), n.d(e, "e", (function() {
                return _
            })), n.d(e, "b", (function() {
                return k
            })), n.d(e, "r", (function() {
                return j
            })), n.d(e, "g", (function() {
                return O
            })), n.d(e, "h", (function() {
                return $
            })), n.d(e, "d", (function() {
                return C
            })), n.d(e, "q", (function() {
                return S
            })), n.d(e, "j", (function() {
                return R
            })), n.d(e, "s", (function() {
                return T
            })), n.d(e, "n", (function() {
                return P
            })), n.d(e, "p", (function() {
                return N
            })), n.d(e, "f", (function() {
                return z
            })), n.d(e, "c", (function() {
                return D
            })), n.d(e, "i", (function() {
                return L
            })), n.d(e, "o", (function() {
                return M
            })), n.d(e, "a", (function() {
                return J
            }));
            n(105), n(53), n(34), n(58), n(93), n(94);
            var r = n(51),
                o = (n(95), n(226), n(227), n(48)),
                c = (n(40), n(41), n(229), n(232), n(159), n(20), n(2)),
                l = (n(50), n(24), n(13), n(52), n(29), n(46)),
                f = n(0);

            function h(object, t) {
                var e = Object.keys(object);
                if (Object.getOwnPropertySymbols) {
                    var n = Object.getOwnPropertySymbols(object);
                    t && (n = n.filter((function(t) {
                        return Object.getOwnPropertyDescriptor(object, t).enumerable
                    }))), e.push.apply(e, n)
                }
                return e
            }
            function d(t) {
                for (var i = 1; i < arguments.length; i++) {
                    var source = null != arguments[i] ? arguments[i] : {};
                    i % 2 ? h(Object(source), !0).forEach((function(e) {
                        Object(l.a)(t, e, source[e])
                    })) : Object.getOwnPropertyDescriptors ? Object.defineProperties(t, Object.getOwnPropertyDescriptors(source)) : h(Object(source)).forEach((function(e) {
                        Object.defineProperty(t, e, Object.getOwnPropertyDescriptor(source, e))
                    }))
                }
                return t
            }
            function m(t, e) {
                var n;
                if ("undefined" == typeof Symbol || null == t[Symbol.iterator]) {
                    if (Array.isArray(t) || (n = function(t, e) {
                        if (!t) return;
                        if ("string" == typeof t) return v(t, e);
                        var n = Object.prototype.toString.call(t).slice(8, -1);
                        "Object" === n && t.constructor && (n = t.constructor.name);
                        if ("Map" === n || "Set" === n) return Array.from(t);
                        if ("Arguments" === n || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)) return v(t, e)
                    }(t)) || e && t && "number" == typeof t.length) {
                        n && (t = n);
                        var i = 0,
                            r = function() {};
                        return {
                            s: r,
                            n: function() {
                                return i >= t.length ? {
                                    done: !0
                                } : {
                                    done: !1,
                                    value: t[i++]
                                }
                            },
                            e: function(t) {
                                throw t
                            },
                            f: r
                        }
                    }
                    throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")
                }
                var o, c = !0,
                    l = !1;
                return {
                    s: function() {
                        n = t[Symbol.iterator]()
                    },
                    n: function() {
                        var t = n.next();
                        return c = t.done, t
                    },
                    e: function(t) {
                        l = !0, o = t
                    },
                    f: function() {
                        try {
                            c || null == n.
                                return ||n.
                            return ()
                        } finally {
                            if (l) throw o
                        }
                    }
                }
            }
            function v(t, e) {
                (null == e || e > t.length) && (e = t.length);
                for (var i = 0, n = new Array(e); i < e; i++) n[i] = t[i];
                return n
            }
            function y(t) {
                f.
                    default.config.errorHandler && f.
                default.config.errorHandler(t)
            }
            function w(t) {
                return t.then((function(t) {
                    return t.
                        default ||t
                }))
            }
            function x(t) {
                return t.$options && "function" == typeof t.$options.fetch && !t.$options.fetch.length
            }
            function _(t) {
                var e, n = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : [],
                    r = t.$children || [],
                    o = m(r);
                try {
                    for (o.s(); !(e = o.n()).done;) {
                        var c = e.value;
                        c.$fetch ? n.push(c) : c.$children && _(c, n)
                    }
                } catch (t) {
                    o.e(t)
                } finally {
                    o.f()
                }
                return n
            }
            function k(t, e) {
                if (e || !t.options.__hasNuxtData) {
                    var n = t.options._originDataFn || t.options.data ||
                        function() {
                            return {}
                        };
                    t.options._originDataFn = n, t.options.data = function() {
                        var data = n.call(this, this);
                        return this.$ssrContext && (e = this.$ssrContext.asyncData[t.cid]), d(d({}, data), e)
                    }, t.options.__hasNuxtData = !0, t._Ctor && t._Ctor.options && (t._Ctor.options.data = t.options.data)
                }
            }
            function j(t) {
                return t.options && t._Ctor === t || (t.options ? (t._Ctor = t, t.extendOptions = t.options) : (t = f.
                default.extend(t))._Ctor = t, !t.options.name && t.options.__file && (t.options.name = t.options.__file)), t
            }
            function O(t) {
                var e = arguments.length > 1 && void 0 !== arguments[1] && arguments[1],
                    n = arguments.length > 2 && void 0 !== arguments[2] ? arguments[2] : "components";
                return Array.prototype.concat.apply([], t.matched.map((function(t, r) {
                    return Object.keys(t[n]).map((function(o) {
                        return e && e.push(r), t[n][o]
                    }))
                })))
            }
            function $(t) {
                var e = arguments.length > 1 && void 0 !== arguments[1] && arguments[1];
                return O(t, e, "instances")
            }
            function C(t, e) {
                return Array.prototype.concat.apply([], t.matched.map((function(t, n) {
                    return Object.keys(t.components).reduce((function(r, o) {
                        return t.components[o] ? r.push(e(t.components[o], t.instances[o], t, o, n)) : delete t.components[o], r
                    }), [])
                })))
            }
            function S(t, e) {
                return Promise.all(C(t, function() {
                    var t = Object(c.a)(regeneratorRuntime.mark((function t(n, r, o, c) {
                        return regeneratorRuntime.wrap((function(t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    if ("function" != typeof n || n.options) {
                                        t.next = 4;
                                        break
                                    }
                                    return t.next = 3, n();
                                case 3:
                                    n = t.sent;
                                case 4:
                                    return o.components[c] = n = j(n), t.abrupt("return", "function" == typeof e ? e(n, r, o, c) : n);
                                case 6:
                                case "end":
                                    return t.stop()
                            }
                        }), t)
                    })));
                    return function(e, n, r, o) {
                        return t.apply(this, arguments)
                    }
                }()))
            }
            function R(t) {
                return E.apply(this, arguments)
            }
            function E() {
                return (E = Object(c.a)(regeneratorRuntime.mark((function t(e) {
                    return regeneratorRuntime.wrap((function(t) {
                        for (;;) switch (t.prev = t.next) {
                            case 0:
                                if (e) {
                                    t.next = 2;
                                    break
                                }
                                return t.abrupt("return");
                            case 2:
                                return t.next = 4, S(e);
                            case 4:
                                return t.abrupt("return", d(d({}, e), {}, {
                                    meta: O(e).map((function(t, n) {
                                        return d(d({}, t.options.meta), (e.matched[n] || {}).meta)
                                    }))
                                }));
                            case 5:
                            case "end":
                                return t.stop()
                        }
                    }), t)
                })))).apply(this, arguments)
            }
            function T(t, e) {
                return A.apply(this, arguments)
            }
            function A() {
                return (A = Object(c.a)(regeneratorRuntime.mark((function t(e, n) {
                    var c, l, f, h;
                    return regeneratorRuntime.wrap((function(t) {
                        for (;;) switch (t.prev = t.next) {
                            case 0:
                                return e.context || (e.context = {
                                    isStatic: !1,
                                    isDev: !1,
                                    isHMR: !1,
                                    app: e,
                                    store: e.store,
                                    payload: n.payload,
                                    error: n.error,
                                    base: "/",
                                    env: {}
                                }, n.req && (e.context.req = n.req), n.res && (e.context.res = n.res), n.ssrContext && (e.context.ssrContext = n.ssrContext), e.context.redirect = function(t, path, n) {
                                    if (t) {
                                        e.context._redirected = !0;
                                        var r = Object(o.a)(path);
                                        if ("number" == typeof t || "undefined" !== r && "object" !== r || (n = path || {}, path = t, r = Object(o.a)(path), t = 302), "object" === r && (path = e.router.resolve(path).route.fullPath), !/(^[.]{1,2}\/)|(^\/(?!\/))/.test(path)) throw path = H(path, n), window.location.replace(path), new Error("ERR_REDIRECT");
                                        e.context.next({
                                            path: path,
                                            query: n,
                                            status: t
                                        })
                                    }
                                }, e.context.nuxtState = window.__NUXT__), t.next = 3, Promise.all([R(n.route), R(n.from)]);
                            case 3:
                                c = t.sent, l = Object(r.a)(c, 2), f = l[0], h = l[1], n.route && (e.context.route = f), n.from && (e.context.from = h), e.context.next = n.next, e.context._redirected = !1, e.context._errored = !1, e.context.isHMR = !1, e.context.params = e.context.route.params || {}, e.context.query = e.context.route.query || {};
                            case 15:
                            case "end":
                                return t.stop()
                        }
                    }), t)
                })))).apply(this, arguments)
            }
            function P(t, e) {
                return !t.length || e._redirected || e._errored ? Promise.resolve() : N(t[0], e).then((function() {
                    return P(t.slice(1), e)
                }))
            }
            function N(t, e) {
                var n;
                return (n = 2 === t.length ? new Promise((function(n) {
                    t(e, (function(t, data) {
                        t && e.error(t), n(data = data || {})
                    }))
                })) : t(e)) && n instanceof Promise && "function" == typeof n.then ? n : Promise.resolve(n)
            }
            function z(base, t) {
                var path = decodeURI(window.location.pathname);
                return "hash" === t ? window.location.hash.replace(/^#\//, "") : (base && 0 === path.indexOf(base) && (path = path.slice(base.length)), (path || "/") + window.location.search + window.location.hash)
            }
            function D(t, e) {
                return function(t, e) {
                    for (var n = new Array(t.length), i = 0; i < t.length; i++)"object" === Object(o.a)(t[i]) && (n[i] = new RegExp("^(?:" + t[i].pattern + ")$", K(e)));
                    return function(e, r) {
                        for (var path = "", data = e || {}, o = (r || {}).pretty ? F : encodeURIComponent, c = 0; c < t.length; c++) {
                            var l = t[c];
                            if ("string" != typeof l) {
                                var f = data[l.name || "pathMatch"],
                                    h = void 0;
                                if (null == f) {
                                    if (l.optional) {
                                        l.partial && (path += l.prefix);
                                        continue
                                    }
                                    throw new TypeError('Expected "' + l.name + '" to be defined')
                                }
                                if (Array.isArray(f)) {
                                    if (!l.repeat) throw new TypeError('Expected "' + l.name + '" to not repeat, but received `' + JSON.stringify(f) + "`");
                                    if (0 === f.length) {
                                        if (l.optional) continue;
                                        throw new TypeError('Expected "' + l.name + '" to not be empty')
                                    }
                                    for (var d = 0; d < f.length; d++) {
                                        if (h = o(f[d]), !n[c].test(h)) throw new TypeError('Expected all "' + l.name + '" to match "' + l.pattern + '", but received `' + JSON.stringify(h) + "`");
                                        path += (0 === d ? l.prefix : l.delimiter) + h
                                    }
                                } else {
                                    if (h = l.asterisk ? F(f, !0) : o(f), !n[c].test(h)) throw new TypeError('Expected "' + l.name + '" to match "' + l.pattern + '", but received "' + h + '"');
                                    path += l.prefix + h
                                }
                            } else path += l
                        }
                        return path
                    }
                }(function(t, e) {
                    var n, r = [],
                        o = 0,
                        c = 0,
                        path = "",
                        l = e && e.delimiter || "/";
                    for (; null != (n = I.exec(t));) {
                        var f = n[0],
                            h = n[1],
                            d = n.index;
                        if (path += t.slice(c, d), c = d + f.length, h) path += h[1];
                        else {
                            var m = t[c],
                                v = n[2],
                                y = n[3],
                                w = n[4],
                                x = n[5],
                                _ = n[6],
                                k = n[7];
                            path && (r.push(path), path = "");
                            var j = null != v && null != m && m !== v,
                                O = "+" === _ || "*" === _,
                                $ = "?" === _ || "*" === _,
                                C = n[2] || l,
                                pattern = w || x;
                            r.push({
                                name: y || o++,
                                prefix: v || "",
                                delimiter: C,
                                optional: $,
                                repeat: O,
                                partial: j,
                                asterisk: Boolean(k),
                                pattern: pattern ? B(pattern) : k ? ".*" : "[^" + U(C) + "]+?"
                            })
                        }
                    }
                    c < t.length && (path += t.substr(c));
                    path && r.push(path);
                    return r
                }(t, e), e)
            }
            function L(t, e) {
                var n = {},
                    r = d(d({}, t), e);
                for (var o in r) String(t[o]) !== String(e[o]) && (n[o] = !0);
                return n
            }
            function M(t) {
                var e;
                if (t.message || "string" == typeof t) e = t.message || t;
                else try {
                    e = JSON.stringify(t, null, 2)
                } catch (n) {
                    e = "[".concat(t.constructor.name, "]")
                }
                return d(d({}, t), {}, {
                    message: e,
                    statusCode: t.statusCode || t.status || t.response && t.response.status || 500
                })
            }
            window.onNuxtReadyCbs = [], window.onNuxtReady = function(t) {
                window.onNuxtReadyCbs.push(t)
            };
            var I = new RegExp(["(\\\\.)", "([\\/.])?(?:(?:\\:(\\w+)(?:\\(((?:\\\\.|[^\\\\()])+)\\))?|\\(((?:\\\\.|[^\\\\()])+)\\))([+*?])?|(\\*))"].join("|"), "g");

            function F(t, e) {
                var n = e ? /[?#]/g : /[/?#]/g;
                return encodeURI(t).replace(n, (function(t) {
                    return "%" + t.charCodeAt(0).toString(16).toUpperCase()
                }))
            }
            function U(t) {
                return t.replace(/([.+*?=^!:${}()[\]|/\\])/g, "\\$1")
            }
            function B(t) {
                return t.replace(/([=!:$/()])/g, "\\$1")
            }
            function K(t) {
                return t && t.sensitive ? "" : "i"
            }
            function H(t, e) {
                var n, o = t.indexOf("://"); - 1 !== o ? (n = t.substring(0, o), t = t.substring(o + 3)) : t.startsWith("//") && (t = t.substring(2));
                var c, l = t.split("/"),
                    f = (n ? n + "://" : "//") + l.shift(),
                    path = l.filter(Boolean).join("/");
                if (2 === (l = path.split("#")).length) {
                    var h = l,
                        d = Object(r.a)(h, 2);
                    path = d[0], c = d[1]
                }
                return f += path ? "/" + path : "", e && "{}" !== JSON.stringify(e) && (f += (2 === t.split("?").length ? "&" : "?") +
                    function(t) {
                        return Object.keys(t).sort().map((function(e) {
                            var n = t[e];
                            return null == n ? "" : Array.isArray(n) ? n.slice().map((function(t) {
                                return [e, "=", t].join("")
                            })).join("&") : e + "=" + n
                        })).filter(Boolean).join("&")
                    }(e)), f += c ? "#" + c : ""
            }
            function J(t, e, n) {
                t.$options[e] || (t.$options[e] = []), t.$options[e].includes(n) || t.$options[e].push(n)
            }
        },
        128: function(t, e, n) {
            "use strict";
            e.a = {}
        },
        131: function(t, e) {
            "serviceWorker" in navigator ? navigator.serviceWorker.register("/sw.js", {
                scope: "/"
            }).then((function(t) {
                window.$sw = t
            })).
            catch ((function(t) {
                console.error("Service worker registration failed:", t)
            })) : console.warn("Service workers are not supported.")
        },
        132: function(t, e, n) {
            n(95), function() {
                var t = document.createElement("script"),
                    e = window.location.protocol.split(":")[0];
                t.src = "https" === e ? "https://zz.bdstatic.com/linksubmit/push.js" : "http://push.zhanzhang.baidu.com/push.js";
                var s = document.getElementsByTagName("script")[0];
                s.parentNode.insertBefore(t, s)
            }()
        },
        133: function(t, e, n) {
            "use strict";
            var r = n(3),
                o = {
                    props: ["links"],
                    data: function() {
                        return {
                            year: (new Date).getFullYear(),
                            project_name: r.p,
                            domain: r.b
                        }
                    }
                },
                c = (n(238), n(7)),
                component = Object(c.a)(o, (function() {
                    var t = this,
                        e = t.$createElement,
                        n = t._self._c || e;
                    return n("footer", {
                        staticClass: "footer"
                    }, [n("div", {
                        staticClass: "copyright-wrap"
                    }, [n("p", [t._v("@2014 - " + t._s(t.year) + " " + t._s(t.project_name) + " " + t._s(t.domain) + " - "), n("nuxt-link", {
                        attrs: {
                            to: "/copyright"
                        }
                    }, [t._v("免责声明")]), t._v(" - "), n("a", {
                        attrs: {
                            href: "http://alibabachanpinjingliwenjuan.mikecrm.com/LbkKKBU",
                            target: "_blank"
                        }
                    }, [t._v("资源举报")]), t._v(" - "), n("a", {
                        attrs: {
                            href: "http://alibabachanpinjingliwenjuan.mikecrm.com/dTNxgJa",
                            target: "_blank"
                        }
                    }, [t._v("广告合作")]), t._v("- "), n("a", {
                        attrs: {
                            href: "http://alibabachanpinjingliwenjuan.mikecrm.com/Mzb3Hpe",
                            target: "_blank"
                        }
                    }, [t._v("建议反馈")])], 1), t._v(" "), n("p", [t._v("以上内容由网络爬虫自动抓取，以非人工方式自动生成。" + t._s(t.project_name) + "不储存、复制、传播任何文件，其网盘资源文件的完整性需要您自行判断。")]), t._v(" "), t.links ? n("p", [n("span", {
                        staticClass: "em"
                    }, [t._v("友链：")]), t._v(" "), n("nuxt-link", {
                        attrs: {
                            to: "/rank"
                        }
                    }, [t._v("资源排行")]), t._v(" "), t._l(t.links, (function(e) {
                        return n("a", {
                            key: e.link,
                            staticClass: "friend-link",
                            attrs: {
                                href: e.link,
                                target: "_blank",
                                rel: "nofollow external noopener"
                            }
                        }, [t._v(t._s(e.text))])
                    }))], 2) : t._e()])])
                }), [], !1, null, "79886f18", null);
            e.a = component.exports
        },
        134: function(t, e, n) {
            "use strict";
            var r = {
                    props: ["keyword"],
                    components: {
                        Search: n(186).a
                    }
                },
                o = (n(237), n(7)),
                component = Object(o.a)(r, (function() {
                    var t = this.$createElement,
                        e = this._self._c || t;
                    return e("div", {
                        staticClass: "header-wrap"
                    }, [e("header", {
                        staticClass: "header"
                    }, [e("div", {
                        staticClass: "header-inner"
                    }, [e("nuxt-link", {
                        attrs: {
                            to: "/"
                        }
                    }, [e("div", {
                        staticClass: "logo-wrap"
                    }, [e("img", {
                        staticClass: "logo",
                        attrs: {
                            src: n(195),
                            alt: "" + this.PROJECT_NAME
                        }
                    }), this._v(" "), e("h1", [this._v(this._s(this.PROJECT_NAME))])])]), this._v(" "), e("div", {
                        staticClass: "form-wrap"
                    }, [e("Search", {
                        attrs: {
                            keyword: this.keyword
                        }
                    })], 1)], 1)])])
                }), [], !1, null, "733dda38", null);
            e.a = component.exports
        },
        186: function(t, e, n) {
            "use strict";
            n(20);
            var r = n(2),
                o = {
                    props: ["keyword"],
                    data: function() {
                        return {
                            searchText: ""
                        }
                    },
                    mounted: function() {
                        var t = this;
                        return Object(r.a)(regeneratorRuntime.mark((function e() {
                            return regeneratorRuntime.wrap((function(e) {
                                for (;;) switch (e.prev = e.next) {
                                    case 0:
                                        t.keyword && (t.searchText = t.keyword);
                                    case 1:
                                    case "end":
                                        return e.stop()
                                }
                            }), e)
                        })))()
                    }
                },
                c = (n(236), n(7)),
                component = Object(c.a)(o, (function() {
                    var t = this,
                        e = t.$createElement,
                        n = t._self._c || e;
                    return n("form", {
                        staticClass: "search-form",
                        attrs: {
                            action: "/search",
                            method: "get"
                        }
                    }, [n("input", {
                        directives: [{
                            name: "model",
                            rawName: "v-model",
                            value: t.searchText,
                            expression: "searchText"
                        }],
                        staticClass: "input",
                        attrs: {
                            type: "text",
                            name: "keyword",
                            placeholder: "请输入你要搜索的关键词",
                            autofocus: ""
                        },
                        domProps: {
                            value: t.searchText
                        },
                        on: {
                            input: function(e) {
                                e.target.composing || (t.searchText = e.target.value)
                            }
                        }
                    }), t._v(" "), n("input", {
                        staticClass: "search-button",
                        attrs: {
                            type: "submit",
                            value: ""
                        }
                    })])
                }), [], !1, null, "10b6ed7a", null);
            e.a = component.exports
        },
        19: function(t, e, n) {
            "use strict";
            n.d(e, "b", (function() {
                return Ot
            })), n.d(e, "a", (function() {
                return T
            }));
            n(20), n(105), n(34), n(24), n(13), n(52);
            var r = n(2),
                o = n(46),
                c = (n(29), n(0)),
                l = n(188),
                f = n(129),
                h = n.n(f),
                d = n(49),
                m = n.n(d),
                v = n(63),
                y = n(1);
            "scrollRestoration" in window.history && (window.history.scrollRestoration = "manual", window.addEventListener("beforeunload", (function() {
                window.history.scrollRestoration = "auto"
            })), window.addEventListener("load", (function() {
                window.history.scrollRestoration = "manual"
            })));
            var w = function() {},
                x = v.a.prototype.push;
            v.a.prototype.push = function(t) {
                var e = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : w,
                    n = arguments.length > 2 ? arguments[2] : void 0;
                return x.call(this, t, e, n)
            }, c.
            default.use(v.a);
            var _ = {
                mode: "history",
                base: decodeURI("/"),
                linkActiveClass: "nuxt-link-active",
                linkExactActiveClass: "nuxt-link-exact-active",
                scrollBehavior: function(t, e, n) {
                    var r = !1,
                        o = Object(y.g)(t);
                    (o.length < 2 && o.every((function(t) {
                        return !1 !== t.options.scrollToTop
                    })) || o.some((function(t) {
                        return t.options.scrollToTop
                    }))) && (r = {
                        x: 0,
                        y: 0
                    }), n && (r = n);
                    var c = window.$nuxt;
                    return t.path === e.path && t.hash !== e.hash && c.$nextTick((function() {
                        return c.$emit("triggerScroll")
                    })), new Promise((function(e) {
                        c.$once("triggerScroll", (function() {
                            if (t.hash) {
                                var n = t.hash;
                                void 0 !== window.CSS && void 0 !== window.CSS.escape && (n = "#" + window.CSS.escape(n.substr(1)));
                                try {
                                    document.querySelector(n) && (r = {
                                        selector: n
                                    })
                                } catch (t) {
                                    console.warn("Failed to save scroll position. Please add CSS.escape() polyfill (https://github.com/mathiasbynens/CSS.escape).")
                                }
                            }
                            e(r)
                        }))
                    }))
                },
                routes: [{
                    path: "/complaint",
                    component: function() {
                        return Object(y.m)(n.e(5).then(n.bind(null, 383)))
                    },
                    name: "complaint"
                }, {
                    path: "/copyright",
                    component: function() {
                        return Object(y.m)(n.e(6).then(n.bind(null, 384)))
                    },
                    name: "copyright"
                }, {
                    path: "/rank",
                    component: function() {
                        return Object(y.m)(n.e(9).then(n.bind(null, 385)))
                    },
                    name: "rank"
                }, {
                    path: "/search",
                    component: function() {
                        return Object(y.m)(Promise.all([n.e(13), n.e(10)]).then(n.bind(null, 381)))
                    },
                    name: "search"
                }, {
                    path: "/about/agree",
                    component: function() {
                        return Object(y.m)(n.e(1).then(n.bind(null, 386)))
                    },
                    name: "about-agree"
                }, {
                    path: "/about/policy",
                    component: function() {
                        return Object(y.m)(n.e(2).then(n.bind(null, 387)))
                    },
                    name: "about-policy"
                }, {
                    path: "/about/safe",
                    component: function() {
                        return Object(y.m)(n.e(3).then(n.bind(null, 388)))
                    },
                    name: "about-safe"
                }, {
                    path: "/about/version",
                    component: function() {
                        return Object(y.m)(n.e(4).then(n.bind(null, 389)))
                    },
                    name: "about-version"
                }, {
                    path: "/detail/:id?",
                    component: function() {
                        return Object(y.m)(n.e(7).then(n.bind(null, 382)))
                    },
                    name: "detail-id"
                }, {
                    path: "/",
                    component: function() {
                        return Object(y.m)(n.e(8).then(n.bind(null, 390)))
                    },
                    name: "index"
                }],
                fallback: !1
            };

            function k() {
                return new v.a(_)
            }
            var j = {
                    name: "NuxtChild",
                    functional: !0,
                    props: {
                        nuxtChildKey: {
                            type: String,
                            default:
                                ""
                        },
                        keepAlive: Boolean,
                        keepAliveProps: {
                            type: Object,
                            default:
                                void 0
                        }
                    },
                    render: function(t, e) {
                        var n = e.parent,
                            data = e.data,
                            r = e.props,
                            o = n.$createElement;
                        data.nuxtChild = !0;
                        for (var c = n, l = n.$nuxt.nuxt.transitions, f = n.$nuxt.nuxt.defaultTransition, h = 0; n;) n.$vnode && n.$vnode.data.nuxtChild && h++, n = n.$parent;
                        data.nuxtChildDepth = h;
                        var d = l[h] || f,
                            m = {};
                        O.forEach((function(t) {
                            void 0 !== d[t] && (m[t] = d[t])
                        }));
                        var v = {};
                        $.forEach((function(t) {
                            "function" == typeof d[t] && (v[t] = d[t].bind(c))
                        }));
                        var y = v.beforeEnter;
                        if (v.beforeEnter = function(t) {
                            if (window.$nuxt.$nextTick((function() {
                                window.$nuxt.$emit("triggerScroll")
                            })), y) return y.call(c, t)
                        }, !1 === d.css) {
                            var w = v.leave;
                            (!w || w.length < 2) && (v.leave = function(t, e) {
                                w && w.call(c, t), c.$nextTick(e)
                            })
                        }
                        var x = o("routerView", data);
                        return r.keepAlive && (x = o("keep-alive", {
                            props: r.keepAliveProps
                        }, [x])), o("transition", {
                            props: m,
                            on: v
                        }, [x])
                    }
                },
                O = ["name", "mode", "appear", "css", "type", "duration", "enterClass", "leaveClass", "appearClass", "enterActiveClass", "enterActiveClass", "leaveActiveClass", "appearActiveClass", "enterToClass", "leaveToClass", "appearToClass"],
                $ = ["beforeEnter", "enter", "afterEnter", "enterCancelled", "beforeLeave", "leave", "afterLeave", "leaveCancelled", "beforeAppear", "appear", "afterAppear", "appearCancelled"],
                C = n(134),
                S = n(133),
                R = (n(193), {
                    props: ["error"],
                    data: function() {
                        return {
                            hot: null,
                            links: null
                        }
                    },
                    components: {
                        HeaderComponent: C.a,
                        FooterComponent: S.a
                    },
                    mounted: function() {
                        return Object(r.a)(regeneratorRuntime.mark((function t() {
                            return regeneratorRuntime.wrap((function(t) {
                                for (;;) switch (t.prev = t.next) {
                                    case 0:
                                    case "end":
                                        return t.stop()
                                }
                            }), t)
                        })))()
                    }
                }),
                E = (n(255), n(7)),
                T = Object(E.a)(R, (function() {
                    var t = this,
                        e = t.$createElement,
                        r = t._self._c || e;
                    return r("div", {
                        staticClass: "error-page"
                    }, [r("header-component"), t._v(" "), r("img", {
                        attrs: {
                            src: n(194),
                            alt: "请求出错"
                        }
                    }), t._v(" "), 404 === t.error.statusCode ? r("h1", {
                        staticClass: "error-tip"
                    }, [t._v("该页面不存在 404:(")]) : r("h1", {
                        staticClass: "error-tip"
                    }, [t._v("该页面好像出了点问题 :(")]), t._v(" "), r("br"), r("br"), r("br"), t._v(" "), r("nuxt-link", {
                        attrs: {
                            to: "/"
                        }
                    }, [t._v("返回首页")]), t._v(" "), r("div", {
                        staticClass: "hot-wrap"
                    }), t._v(" "), r("footer-component")], 1)
                }), [], !1, null, null, null).exports,
                A = (n(40), n(41), n(159), n(51)),
                P = {
                    name: "Nuxt",
                    components: {
                        NuxtChild: j,
                        NuxtError: T
                    },
                    props: {
                        nuxtChildKey: {
                            type: String,
                            default:
                                void 0
                        },
                        keepAlive: Boolean,
                        keepAliveProps: {
                            type: Object,
                            default:
                                void 0
                        },
                        name: {
                            type: String,
                            default:
                                "default"
                        }
                    },
                    errorCaptured: function(t) {
                        this.displayingNuxtError && (this.errorFromNuxtError = t, this.$forceUpdate())
                    },
                    computed: {
                        routerViewKey: function() {
                            if (void 0 !== this.nuxtChildKey || this.$route.matched.length > 1) return this.nuxtChildKey || Object(y.c)(this.$route.matched[0].path)(this.$route.params);
                            var t = Object(A.a)(this.$route.matched, 1)[0];
                            if (!t) return this.$route.path;
                            var e = t.components.
                                default;
                            if (e && e.options) {
                                var n = e.options;
                                if (n.key) return "function" == typeof n.key ? n.key(this.$route) : n.key
                            }
                            return /\/$/.test(t.path) ? this.$route.path:
                                this.$route.path.replace(/\/$/, "")
                        }
                    },
                    beforeCreate: function() {
                        c.
                        default.util.defineReactive(this, "nuxt", this.$root.$options.nuxt)
                    },
                    render: function(t) {
                        var e = this;
                        return this.nuxt.err ? this.errorFromNuxtError ? (this.$nextTick((function() {
                            return e.errorFromNuxtError = !1
                        })), t("div", {}, [t("h2", "An error occured while showing the error page"), t("p", "Unfortunately an error occured and while showing the error page another error occured"), t("p", "Error details: ".concat(this.errorFromNuxtError.toString())), t("nuxt-link", {
                            props: {
                                to: "/"
                            }
                        }, "Go back to home")])) : (this.displayingNuxtError = !0, this.$nextTick((function() {
                            return e.displayingNuxtError = !1
                        })), t(T, {
                            props: {
                                error: this.nuxt.err
                            }
                        })) : t("NuxtChild", {
                            key: this.routerViewKey,
                            props: this.$props
                        })
                    }
                },
                N = (n(53), n(58), n(50), {
                    name: "NuxtLoading",
                    data: function() {
                        return {
                            percent: 0,
                            show: !1,
                            canSucceed: !0,
                            reversed: !1,
                            skipTimerCount: 0,
                            rtl: !1,
                            throttle: 200,
                            duration: 5e3,
                            continuous: !1
                        }
                    },
                    computed: {
                        left: function() {
                            return !(!this.continuous && !this.rtl) && (this.rtl ? this.reversed ? "0px" : "auto" : this.reversed ? "auto" : "0px")
                        }
                    },
                    beforeDestroy: function() {
                        this.clear()
                    },
                    methods: {
                        clear: function() {
                            clearInterval(this._timer), clearTimeout(this._throttle), this._timer = null
                        },
                        start: function() {
                            var t = this;
                            return this.clear(), this.percent = 0, this.reversed = !1, this.skipTimerCount = 0, this.canSucceed = !0, this.throttle ? this._throttle = setTimeout((function() {
                                return t.startTimer()
                            }), this.throttle) : this.startTimer(), this
                        },
                        set: function(t) {
                            return this.show = !0, this.canSucceed = !0, this.percent = Math.min(100, Math.max(0, Math.floor(t))), this
                        },
                        get: function() {
                            return this.percent
                        },
                        increase: function(t) {
                            return this.percent = Math.min(100, Math.floor(this.percent + t)), this
                        },
                        decrease: function(t) {
                            return this.percent = Math.max(0, Math.floor(this.percent - t)), this
                        },
                        pause: function() {
                            return clearInterval(this._timer), this
                        },
                        resume: function() {
                            return this.startTimer(), this
                        },
                        finish: function() {
                            return this.percent = this.reversed ? 0 : 100, this.hide(), this
                        },
                        hide: function() {
                            var t = this;
                            return this.clear(), setTimeout((function() {
                                t.show = !1, t.$nextTick((function() {
                                    t.percent = 0, t.reversed = !1
                                }))
                            }), 500), this
                        },
                        fail: function(t) {
                            return this.canSucceed = !1, this
                        },
                        startTimer: function() {
                            var t = this;
                            this.show || (this.show = !0), void 0 === this._cut && (this._cut = 1e4 / Math.floor(this.duration)), this._timer = setInterval((function() {
                                t.skipTimerCount > 0 ? t.skipTimerCount-- : (t.reversed ? t.decrease(t._cut) : t.increase(t._cut), t.continuous && (t.percent >= 100 || t.percent <= 0) && (t.skipTimerCount = 1, t.reversed = !t.reversed))
                            }), 100)
                        }
                    },
                    render: function(t) {
                        var e = t(!1);
                        return this.show && (e = t("div", {
                            staticClass: "nuxt-progress",
                            class: {
                                "nuxt-progress-notransition": this.skipTimerCount > 0,
                                "nuxt-progress-failed": !this.canSucceed
                            },
                            style: {
                                width: this.percent + "%",
                                left: this.left
                            }
                        })), e
                    }
                }),
                z = (n(256), Object(E.a)(N, void 0, void 0, !1, null, null, null).exports),
                D = (n(257), [function() {
                    var t = this.$createElement,
                        e = this._self._c || t;
                    return e("div", {
                        staticClass: "feedback-wrap"
                    }, [e("div", {
                        staticClass: "feedback-item"
                    }, [e("img", {
                        staticClass: "img",
                        attrs: {
                            src: n(197),
                            alt: "微信公众号"
                        }
                    }), this._v(" "), e("div", {
                        staticClass: "hover-tip"
                    }, [e("p", {
                        staticClass: "text"
                    }, [this._v("微信扫码反馈问题")]), this._v(" "), e("img", {
                        attrs: {
                            src: n(258),
                            alt: "微信公众号二维码"
                        }
                    })])]), this._v(" "), e("a", {
                        staticClass: "feedback-item",
                        attrs: {
                            href: "http://alibabachanpinjingliwenjuan.mikecrm.com/Mzb3Hpe",
                            target: "_blank"
                        }
                    }, [e("img", {
                        staticClass: "img",
                        attrs: {
                            src: n(198),
                            alt: "在线反馈"
                        }
                    }), this._v(" "), e("div", {
                        staticClass: "hover-tip"
                    }, [e("p", [e("a", {
                        attrs: {
                            href: "http://alibabachanpinjingliwenjuan.mikecrm.com/Mzb3Hpe",
                            target: "_blank"
                        }
                    }, [this._v("点此在线反馈问题")])])])])])
                }]),
                L = {},
                M = (n(259), {
                    methods: {
                        hideFilter: function() {
                            this.$store.commit("hideFilter")
                        }
                    },
                    components: {
                        Feedback: Object(E.a)(L, (function() {
                            var t = this.$createElement;
                            this._self._c;
                            return this._m(0)
                        }), D, !1, null, "12bcd8a1", null).exports
                    }
                }),
                I = (n(260), Object(E.a)(M, (function() {
                    var t = this.$createElement,
                        e = this._self._c || t;
                    return e("div", {
                        staticClass: "nuxt-wrap",
                        on: {
                            click: this.hideFilter
                        }
                    }, [e("nuxt"), this._v(" "), e("Feedback")], 1)
                }), [], !1, null, null, null).exports);

            function F(t, e) {
                var n;
                if ("undefined" == typeof Symbol || null == t[Symbol.iterator]) {
                    if (Array.isArray(t) || (n = function(t, e) {
                        if (!t) return;
                        if ("string" == typeof t) return U(t, e);
                        var n = Object.prototype.toString.call(t).slice(8, -1);
                        "Object" === n && t.constructor && (n = t.constructor.name);
                        if ("Map" === n || "Set" === n) return Array.from(t);
                        if ("Arguments" === n || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)) return U(t, e)
                    }(t)) || e && t && "number" == typeof t.length) {
                        n && (t = n);
                        var i = 0,
                            r = function() {};
                        return {
                            s: r,
                            n: function() {
                                return i >= t.length ? {
                                    done: !0
                                } : {
                                    done: !1,
                                    value: t[i++]
                                }
                            },
                            e: function(t) {
                                throw t
                            },
                            f: r
                        }
                    }
                    throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")
                }
                var o, c = !0,
                    l = !1;
                return {
                    s: function() {
                        n = t[Symbol.iterator]()
                    },
                    n: function() {
                        var t = n.next();
                        return c = t.done, t
                    },
                    e: function(t) {
                        l = !0, o = t
                    },
                    f: function() {
                        try {
                            c || null == n.
                                return ||n.
                            return ()
                        } finally {
                            if (l) throw o
                        }
                    }
                }
            }
            function U(t, e) {
                (null == e || e > t.length) && (e = t.length);
                for (var i = 0, n = new Array(e); i < e; i++) n[i] = t[i];
                return n
            }
            var B = {
                    _default: Object(y.r)(I)
                },
                K = {
                    render: function(t, e) {
                        var n = t("NuxtLoading", {
                            ref: "loading"
                        });
                        if (this.nuxt.err && T) {
                            var r = (T.options || T).layout;
                            r && this.setLayout("function" == typeof r ? r.call(T, this.context) : r)
                        }
                        var o = t(this.layout || "nuxt"),
                            c = t("div", {
                                domProps: {
                                    id: "__layout"
                                },
                                key: this.layoutName
                            }, [o]),
                            l = t("transition", {
                                props: {
                                    name: "layout",
                                    mode: "out-in"
                                },
                                on: {
                                    beforeEnter: function(t) {
                                        window.$nuxt.$nextTick((function() {
                                            window.$nuxt.$emit("triggerScroll")
                                        }))
                                    }
                                }
                            }, [c]);
                        return t("div", {
                            domProps: {
                                id: "__nuxt"
                            }
                        }, [n, l])
                    },
                    data: function() {
                        return {
                            isOnline: !0,
                            layout: null,
                            layoutName: "",
                            nbFetching: 0
                        }
                    },
                    beforeCreate: function() {
                        c.
                        default.util.defineReactive(this, "nuxt", this.$options.nuxt)
                    },
                    created: function() {
                        c.
                            default.prototype.$nuxt = this, window.$nuxt = this, this.refreshOnlineStatus(), window.addEventListener("online", this.refreshOnlineStatus), window.addEventListener("offline", this.refreshOnlineStatus), this.error = this.nuxt.error, this.context = this.$options.context
                    },
                    mounted: function() {
                        this.$loading = this.$refs.loading
                    },
                    watch: {
                        "nuxt.err": "errorChanged"
                    },
                    computed: {
                        isOffline: function() {
                            return !this.isOnline
                        },
                        isFetching: function() {
                            return this.nbFetching > 0
                        }
                    },
                    methods: {
                        refreshOnlineStatus: function() {
                            void 0 === window.navigator.onLine ? this.isOnline = !0 : this.isOnline = window.navigator.onLine
                        },
                        refresh: function() {
                            var t = this;
                            return Object(r.a)(regeneratorRuntime.mark((function e() {
                                var n, r;
                                return regeneratorRuntime.wrap((function(e) {
                                    for (;;) switch (e.prev = e.next) {
                                        case 0:
                                            if ((n = Object(y.h)(t.$route)).length) {
                                                e.next = 3;
                                                break
                                            }
                                            return e.abrupt("return");
                                        case 3:
                                            return t.$loading.start(), r = n.map((function(e) {
                                                var p = [];
                                                if (e.$options.fetch && e.$options.fetch.length && p.push(Object(y.p)(e.$options.fetch, t.context)), e.$fetch) p.push(e.$fetch());
                                                else {
                                                    var n, r = F(Object(y.e)(e.$vnode.componentInstance));
                                                    try {
                                                        for (r.s(); !(n = r.n()).done;) {
                                                            var component = n.value;
                                                            p.push(component.$fetch())
                                                        }
                                                    } catch (t) {
                                                        r.e(t)
                                                    } finally {
                                                        r.f()
                                                    }
                                                }
                                                return e.$options.asyncData && p.push(Object(y.p)(e.$options.asyncData, t.context).then((function(t) {
                                                    for (var n in t) c.
                                                    default.set(e.$data, n, t[n])
                                                }))), Promise.all(p)
                                            })), e.prev = 5, e.next = 8, Promise.all(r);
                                        case 8:
                                            e.next = 15;
                                            break;
                                        case 10:
                                            e.prev = 10, e.t0 = e.
                                            catch (5), t.$loading.fail(e.t0), Object(y.k)(e.t0), t.error(e.t0);
                                        case 15:
                                            t.$loading.finish();
                                        case 16:
                                        case "end":
                                            return e.stop()
                                    }
                                }), e, null, [
                                    [5, 10]
                                ])
                            })))()
                        },
                        errorChanged: function() {
                            this.nuxt.err && this.$loading && (this.$loading.fail && this.$loading.fail(this.nuxt.err), this.$loading.finish && this.$loading.finish())
                        },
                        setLayout: function(t) {
                            return t && B["_" + t] || (t = "default"), this.layoutName = t, this.layout = B["_" + t], this.layout
                        },
                        loadLayout: function(t) {
                            return t && B["_" + t] || (t = "default"), Promise.resolve(B["_" + t])
                        }
                    },
                    components: {
                        NuxtLoading: z
                    }
                },
                H = n(130);
            c.
            default.use(H.a);
            var J = {};
            (J = function(t, e) {
                if ((t = t.
                    default ||t).commit) throw new Error("[nuxt] ".concat(e, " should export a method that returns a Vuex instance."));
                return "function" != typeof t && (t = Object.assign({}, t)), function(t, e) {
                    if (t.state && "function" != typeof t.state) {
                        console.warn("'state' should be a method that returns an object in ".concat(e));
                        var n = Object.assign({}, t.state);
                        t = Object.assign({}, t, {
                            state: function() {
                                return n
                            }
                        })
                    }
                    return t
                }(t, e)
            }(n(261), "store/index.js")).modules = J.modules || {};
            var Q = J instanceof Function ? J:


                function() {
                    return new H.a.Store(Object.assign({
                        strict: !1
                    }, J))
                };
            var X = n(131),
                V = n.n(X),
                G = {
                    render: function(t) {
                        return t("ins", {
                            class: ["adsbygoogle"],
                            style: this.adStyle,
                            attrs: {
                                "data-ad-client": this.adClient,
                                "data-ad-slot": this.adSlot || null,
                                "data-ad-format": this.adFormat,
                                "data-ad-region": this.show ? this.adRegion() : null,
                                "data-ad-layout": this.adLayout || null,
                                "data-ad-layout-key": this.adLayoutKey || null,
                                "data-page-url": this.pageUrl ? this.pageUrl : null,
                                "data-analytics-uacct": this.analyticsUacct ? this.analyticsUacct : null,
                                "data-analytics-domain-name": this.analyticsDomainName ? this.analyticsDomainName : null,
                                "data-adtest": null,
                                "data-adsbygoogle-status": this.show ? null : ""
                            },
                            domProps: {
                                innerHTML: this.show ? "" : " "
                            },
                            key: Math.random()
                        })
                    },
                    props: {
                        adClient: {
                            type: String,
                            default:
                                "ca-pub-2135749442908535"
                        },
                        adSlot: {
                            type: String
                        },
                        adFormat: {
                            type: String,
                            default:
                                "auto"
                        },
                        adLayout: {
                            type: String
                        },
                        adLayoutKey: {
                            type: String
                        },
                        adStyle: {
                            type: Object,
                            default:


                                function() {
                                    return {
                                        display: "block"
                                    }
                                }
                        },
                        pageUrl: {
                            type: String
                        },
                        analyticsUacct: {
                            type: String,
                            default:
                                ""
                        },
                        analyticsDomainName: {
                            type: String,
                            default:
                                ""
                        },
                        includeQuery: {
                            type: Boolean,
                            default:
                                !1
                        }
                    },
                    data: function() {
                        return {
                            show: !0
                        }
                    },
                    mounted: function() {
                        this.showAd()
                    },
                    watch: {
                        $route: function(t, e) {
                            if (t.fullPath !== e.fullPath) {
                                var n = Object.keys,
                                    r = t.query,
                                    o = e.query,
                                    c = !1;
                                t.path !== e.path ? c = !0 : this.includeQuery && (c = n(r).length !== n(o).length || !n(r).every((function(t) {
                                    return r[t] === o[t]
                                }))), c && this.updateAd()
                            }
                        }
                    },
                    methods: {
                        adRegion: function() {
                            return "page-" + Math.random()
                        },
                        updateAd: function() {
                            this.isServer || (this.show = !1, this.$nextTick(this.showAd))
                        },
                        showAd: function() {
                            this.show = !0, this.$nextTick((function() {
                                try {
                                    (window.adsbygoogle = window.adsbygoogle || []).push({})
                                } catch (t) {
                                    console.error(t)
                                }
                            }))
                        }
                    }
                };
            c.
            default.component("adsbygoogle", G);
            var W = n(47),
                Y = n.n(W),
                Z = n(190),
                tt = n.n(Z);

            function et(t, e) {
                var n;
                if ("undefined" == typeof Symbol || null == t[Symbol.iterator]) {
                    if (Array.isArray(t) || (n = function(t, e) {
                        if (!t) return;
                        if ("string" == typeof t) return nt(t, e);
                        var n = Object.prototype.toString.call(t).slice(8, -1);
                        "Object" === n && t.constructor && (n = t.constructor.name);
                        if ("Map" === n || "Set" === n) return Array.from(t);
                        if ("Arguments" === n || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)) return nt(t, e)
                    }(t)) || e && t && "number" == typeof t.length) {
                        n && (t = n);
                        var i = 0,
                            r = function() {};
                        return {
                            s: r,
                            n: function() {
                                return i >= t.length ? {
                                    done: !0
                                } : {
                                    done: !1,
                                    value: t[i++]
                                }
                            },
                            e: function(t) {
                                throw t
                            },
                            f: r
                        }
                    }
                    throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")
                }
                var o, c = !0,
                    l = !1;
                return {
                    s: function() {
                        n = t[Symbol.iterator]()
                    },
                    n: function() {
                        var t = n.next();
                        return c = t.done, t
                    },
                    e: function(t) {
                        l = !0, o = t
                    },
                    f: function() {
                        try {
                            c || null == n.
                                return ||n.
                            return ()
                        } finally {
                            if (l) throw o
                        }
                    }
                }
            }
            function nt(t, e) {
                (null == e || e > t.length) && (e = t.length);
                for (var i = 0, n = new Array(e); i < e; i++) n[i] = t[i];
                return n
            }
            for (var at = {
                setBaseURL: function(t) {
                    this.defaults.baseURL = t
                },
                setHeader: function(t, e) {
                    var n, r = arguments.length > 2 && void 0 !== arguments[2] ? arguments[2] : "common",
                        o = et(Array.isArray(r) ? r : [r]);
                    try {
                        for (o.s(); !(n = o.n()).done;) {
                            var c = n.value;
                            if (!e) return void delete this.defaults.headers[c][t];
                            this.defaults.headers[c][t] = e
                        }
                    } catch (t) {
                        o.e(t)
                    } finally {
                        o.f()
                    }
                },
                setToken: function(t, e) {
                    var n = arguments.length > 2 && void 0 !== arguments[2] ? arguments[2] : "common",
                        r = t ? (e ? e + " " : "") + t : null;
                    this.setHeader("Authorization", r, n)
                },
                onRequest: function(t) {
                    this.interceptors.request.use((function(e) {
                        return t(e) || e
                    }))
                },
                onResponse: function(t) {
                    this.interceptors.response.use((function(e) {
                        return t(e) || e
                    }))
                },
                onRequestError: function(t) {
                    this.interceptors.request.use(void 0, (function(e) {
                        return t(e) || Promise.reject(e)
                    }))
                },
                onResponseError: function(t) {
                    this.interceptors.response.use(void 0, (function(e) {
                        return t(e) || Promise.reject(e)
                    }))
                },
                onError: function(t) {
                    this.onRequestError(t), this.onResponseError(t)
                },
                create: function(t) {
                    return ft(tt()(t, this.defaults))
                }
            }, ot = function() {
                var t = st[it];
                at["$" + t] = function() {
                    return this[t].apply(this, arguments).then((function(t) {
                        return t && t.data
                    }))
                }
            }, it = 0, st = ["request", "delete", "get", "head", "options", "post", "put", "patch"]; it < st.length; it++) ot();
            var i, s, ut, ct, a, lt, ft = function(t) {
                    var e = Y.a.create(t);
                    return e.CancelToken = Y.a.CancelToken, e.isCancel = Y.a.isCancel, function(t) {
                        for (var e in at) t[e] = at[e].bind(t)
                    }(e), pt(e), e
                },
                pt = function(t) {
                    var e = {
                            finish: function() {},
                            start: function() {},
                            fail: function() {},
                            set: function() {}
                        },
                        n = function() {
                            var t = "undefined" != typeof window && window.$nuxt;
                            return t && t.$loading && t.$loading.set ? t.$loading : e
                        },
                        r = 0;
                    t.onRequest((function(t) {
                        t && !1 === t.progress || r++
                    })), t.onResponse((function(t) {
                        t && t.config && !1 === t.config.progress || --r <= 0 && (r = 0, n().finish())
                    })), t.onError((function(t) {
                        t && t.config && !1 === t.config.progress || (r--, Y.a.isCancel(t) || (n().fail(), n().finish()))
                    }));
                    var o = function(t) {
                        if (r) {
                            var progress = 100 * t.loaded / (t.total * r);
                            n().set(Math.min(100, progress))
                        }
                    };
                    t.defaults.onUploadProgress = o, t.defaults.onDownloadProgress = o
                },
                ht = function(t, e) {
                    var n = ft({
                        baseURL: "http://localhost:3000/",
                        headers: {
                            common: {
                                Accept: "application/json, text/plain, */*"
                            },
                            delete: {},
                            get: {},
                            head: {},
                            post: {},
                            put: {},
                            patch: {}
                        }
                    });
                    t.$axios = n, e("axios", n)
                };
            try {
                i = window, s = document, ut = "script", ct = "ga", i.GoogleAnalyticsObject = ct, i.ga = i.ga ||
                    function() {
                        (i.ga.q = i.ga.q || []).push(arguments)
                    }, i.ga.l = 1 * new Date, a = s.createElement(ut), lt = s.getElementsByTagName(ut)[0], a.async = 1, a.src = "https://www.google-analytics.com/analytics.js", lt.parentNode.insertBefore(a, lt), ga("create", "UA-128641089-12", "auto")
            } catch (t) {}
            var mt = function(t) {
                    t.app.router.afterEach((function(t, e) {
                        try {
                            window._hmt = window._hmt || [], window._hmt.push(["_trackPageview", t.fullPath]), ga("set", "page", t.fullPath), ga("send", "pageview")
                        } catch (t) {}
                    }))
                },
                vt = (n(95), n(3));
            c.
            default.mixin({
                data: function() {
                    return {
                        DOMAIN: vt.b,
                        PROJECT_NAME: vt.p,
                        TITLE: vt.r,
                        DESCRIPTION: vt.a,
                        KEYWORDS: vt.n
                    }
                }
            });
            var yt = n(132),
                gt = n.n(yt),
                bt = function(t) {
                    var e = t.app,
                        n = t.store;
                    e.router.beforeEach((function(t, e, r) {
                        var filter = n.state.filter,
                            o = t.query;
                        "search" == t.name && (-1 !== filter.time || -1 !== filter.size || -1 !== filter.type || -1 !== filter.searchtype) ? void 0 !== o.size || void 0 !== o.time || void 0 !== o.type || void 0 !== o.searchtype ? r() : (-1 !== filter.size && (o.size = filter.size), -1 !== filter.time && (o.time = filter.time), -1 !== filter.type && (o.type = filter.type), -1 !== filter.searchtype && (o.searchtype = filter.searchtype), r({
                            path: "search",
                            query: o
                        })) : r()
                    }))
                },
                wt = n(191),
                xt = n.n(wt);

            function _t(object, t) {
                var e = Object.keys(object);
                if (Object.getOwnPropertySymbols) {
                    var n = Object.getOwnPropertySymbols(object);
                    t && (n = n.filter((function(t) {
                        return Object.getOwnPropertyDescriptor(object, t).enumerable
                    }))), e.push.apply(e, n)
                }
                return e
            }
            function kt(t) {
                for (var i = 1; i < arguments.length; i++) {
                    var source = null != arguments[i] ? arguments[i] : {};
                    i % 2 ? _t(Object(source), !0).forEach((function(e) {
                        Object(o.a)(t, e, source[e])
                    })) : Object.getOwnPropertyDescriptors ? Object.defineProperties(t, Object.getOwnPropertyDescriptors(source)) : _t(Object(source)).forEach((function(e) {
                        Object.defineProperty(t, e, Object.getOwnPropertyDescriptor(source, e))
                    }))
                }
                return t
            }
            c.
            default.use(xt.a), c.
            default.component(h.a.name, h.a), c.
            default.component(m.a.name, kt(kt({}, m.a), {}, {
                render: function(t, e) {
                    return m.a._warned || (m.a._warned = !0, console.warn("<no-ssr> has been deprecated and will be removed in Nuxt 3, please use <client-only> instead")), m.a.render(t, e)
                }
            })), c.
            default.component(j.name, j), c.
            default.component("NChild", j), c.
            default.component(P.name, P), c.
            default.use(l.a, {
                keyName: "head",
                attribute: "data-n-head",
                ssrAttribute: "data-n-head-ssr",
                tagIDKeyName: "hid"
            });
            var jt = {
                name: "page",
                mode: "out-in",
                appear: !1,
                appearClass: "appear",
                appearActiveClass: "appear-active",
                appearToClass: "appear-to"
            };

            function Ot(t) {
                return $t.apply(this, arguments)
            }
            function $t() {
                return ($t = Object(r.a)(regeneratorRuntime.mark((function t(e) {
                    var n, r, o, l, f, h, path, d;
                    return regeneratorRuntime.wrap((function(t) {
                        for (;;) switch (t.prev = t.next) {
                            case 0:
                                return t.next = 2, k();
                            case 2:
                                return n = t.sent, (r = Q(e)).$router = n, o = r.registerModule, r.registerModule = function(path, t, e) {
                                    return o.call(r, path, t, Object.assign({
                                        preserveState: !0
                                    }, e))
                                }, l = kt({
                                    head: {
                                        title: "网盘搜索，就用大力盘搜索 - 最好用的百度网盘搜索引擎 https://www.dalipan.com",
                                        meta: [{
                                            charset: "utf-8"
                                        }, {
                                            name: "baidu_union_verify",
                                            content: "b16e47d4cde7293c80c8012817a3d24c"
                                        }, {
                                            name: "viewport",
                                            content: "width=device-width, initial-scale=1, minimum-scale=1, maximum-scale=1, user-scalable=no"
                                        }, {
                                            hid: "og:site_name",
                                            name: "og:site_name",
                                            property: "og:site_name",
                                            content: "网盘搜索，就用大力盘搜索 - 最好用的百度网盘搜索引擎 https://www.dalipan.com"
                                        }, {
                                            hid: "description",
                                            name: "description",
                                            content: "大力盘搜索支持百度云搜索，可快速搜索百度网盘资源中的有效连接，自动识别无效的百度云网盘资源，每天更新海量资源。"
                                        }, {
                                            hid: "keywords",
                                            name: "keywords",
                                            content: "大力盘搜索,盘搜搜,大力盘,大力盘搜索,大力搜索盘,网盘搜索,电影下载,迅雷下载,bt下载,种子下载,电子书下载,百度云盘搜索,网盘搜索引擎,百度网盘搜索"
                                        }, {
                                            "http-equiv": "Cache-control",
                                            content: "no-transform"
                                        }, {
                                            "http-equiv": "Cache-control",
                                            content: "no-siteapp"
                                        }, {
                                            hid: "mobile-web-app-capable",
                                            name: "mobile-web-app-capable",
                                            content: "yes"
                                        }, {
                                            hid: "apple-mobile-web-app-title",
                                            name: "apple-mobile-web-app-title",
                                            content: "dalipan.com"
                                        }, {
                                            hid: "author",
                                            name: "author",
                                            content: "ChangMM"
                                        }, {
                                            hid: "theme-color",
                                            name: "theme-color",
                                            content: "#FA541C"
                                        }, {
                                            hid: "og:type",
                                            name: "og:type",
                                            property: "og:type",
                                            content: "website"
                                        }, {
                                            hid: "og:title",
                                            name: "og:title",
                                            property: "og:title",
                                            content: "dalipan.com"
                                        }, {
                                            hid: "og:description",
                                            name: "og:description",
                                            property: "og:description",
                                            content: "Baidu network disk resources search engine."
                                        }],
                                        link: [{
                                            rel: "canonical",
                                            href: "https://www.dalipan.com"
                                        }, {
                                            rel: "icon",
                                            type: "image/x-icon",
                                            href: "/favicon.ico"
                                        }, {
                                            rel: "manifest",
                                            href: "/_nuxt/manifest.364331e8.json"
                                        }],
                                        script: [{
                                            src: "https://hm.baidu.com/hm.js?a38cbf40e7912b2414f8dfea35fb4eb9"
                                        }, {
                                            src: "https://jspassport.ssl.qhimg.com/11.0.1.js?d182b3f28525f2db83acfaaf6e696dba",
                                            id: "sozz"
                                        }, {
                                            async: !0,
                                            src: "//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"
                                        }, {
                                            innerHTML: '\n      (adsbygoogle = window.adsbygoogle || []).push({\n        google_ad_client: "ca-pub-2135749442908535",\n        enable_page_level_ads: false\n      });\n  '
                                        }],
                                        style: [],
                                        __dangerouslyDisableSanitizers: ["script"],
                                        htmlAttrs: {
                                            lang: "en"
                                        }
                                    },
                                    store: r,
                                    router: n,
                                    nuxt: {
                                        defaultTransition: jt,
                                        transitions: [jt],
                                        setTransitions: function(t) {
                                            return Array.isArray(t) || (t = [t]), t = t.map((function(t) {
                                                return t = t ? "string" == typeof t ? Object.assign({}, jt, {
                                                    name: t
                                                }) : Object.assign({}, jt, t) : jt
                                            })), this.$options.nuxt.transitions = t, t
                                        },
                                        err: null,
                                        dateErr: null,
                                        error: function(t) {
                                            t = t || null, l.context._errored = Boolean(t), t = t ? Object(y.o)(t) : null;
                                            var n = l.nuxt;
                                            return this && (n = this.nuxt || this.$options.nuxt), n.dateErr = Date.now(), n.err = t, e && (e.nuxt.error = t), t
                                        }
                                    }
                                }, K), r.app = l, f = e ? e.next : function(t) {
                                    return l.router.push(t)
                                }, e ? h = n.resolve(e.url).route : (path = Object(y.f)(n.options.base, n.options.mode), h = n.resolve(path).route), t.next = 13, Object(y.s)(l, {
                                    store: r,
                                    route: h,
                                    next: f,
                                    error: l.nuxt.error.bind(l),
                                    payload: e ? e.payload : void 0,
                                    req: e ? e.req : void 0,
                                    res: e ? e.res : void 0,
                                    beforeRenderFns: e ? e.beforeRenderFns : void 0,
                                    ssrContext: e
                                });
                            case 13:
                                if (d = function(t, e) {
                                    if (!t) throw new Error("inject(key, value) has no key provided");
                                    if (void 0 === e) throw new Error("inject('".concat(t, "', value) has no value provided"));
                                    l[t = "$" + t] = e, r[t] = l[t];
                                    var n = "__nuxt_" + t + "_installed__";
                                    c.
                                        default [n] || (c.
                                        default [n] = !0, c.
                                    default.use((function() {
                                        Object.prototype.hasOwnProperty.call(c.
                                            default, t) || Object.defineProperty(c.
                                            default.prototype, t, {
                                            get: function() {
                                                return this.$root.$options[t]
                                            }
                                        })
                                    })))
                                }, window.__NUXT__ && window.__NUXT__.state && r.replaceState(window.__NUXT__.state), "function" != typeof V.a) {
                                    t.next = 18;
                                    break
                                }
                                return t.next = 18, V()(l.context, d);
                            case 18:
                                t.next = 21;
                                break;
                            case 21:
                                return t.next = 24, ht(l.context, d);
                            case 24:
                                return t.next = 27, mt(l.context);
                            case 27:
                                return t.next = 30, void l.context;
                            case 30:
                                t.next = 33;
                                break;
                            case 33:
                                if ("function" != typeof gt.a) {
                                    t.next = 36;
                                    break
                                }
                                return t.next = 36, gt()(l.context, d);
                            case 36:
                                return t.next = 39, bt(l.context);
                            case 39:
                                t.next = 42;
                                break;
                            case 42:
                                t.next = 45;
                                break;
                            case 45:
                                return t.abrupt("return", {
                                    store: r,
                                    app: l,
                                    router: n
                                });
                            case 46:
                            case "end":
                                return t.stop()
                        }
                    }), t)
                })))).apply(this, arguments)
            }
        },
        192: function(t, e, n) {
            "use strict";
            n(13), n(20);
            var r = n(2),
                o = n(0),
                c = n(1),
                l = window.__NUXT__;

            function f() {
                if (!this._hydrated) return this.$fetch()
            }
            function h() {
                if ((t = this).$vnode && t.$vnode.elm && t.$vnode.elm.dataset && t.$vnode.elm.dataset.fetchKey) {
                    var t;
                    this._hydrated = !0, this._fetchKey = +this.$vnode.elm.dataset.fetchKey;
                    var data = l.fetch[this._fetchKey];
                    if (data && data._error) this.$fetchState.error = data._error;
                    else for (var e in data) o.
                    default.set(this.$data, e, data[e])
                }
            }
            function d() {
                var t = this;
                return this._fetchPromise || (this._fetchPromise = m.call(this).then((function() {
                    delete t._fetchPromise
                }))), this._fetchPromise
            }
            function m() {
                return v.apply(this, arguments)
            }
            function v() {
                return (v = Object(r.a)(regeneratorRuntime.mark((function t() {
                    var e, n, r, o = this;
                    return regeneratorRuntime.wrap((function(t) {
                        for (;;) switch (t.prev = t.next) {
                            case 0:
                                return this.$nuxt.nbFetching++, this.$fetchState.pending = !0, this.$fetchState.error = null, this._hydrated = !1, e = null, n = Date.now(), t.prev = 6, t.next = 9, this.$options.fetch.call(this);
                            case 9:
                                t.next = 14;
                                break;
                            case 11:
                                t.prev = 11, t.t0 = t.
                                catch (6), e = Object(c.o)(t.t0);
                            case 14:
                                if (!((r = this._fetchDelay - (Date.now() - n)) > 0)) {
                                    t.next = 18;
                                    break
                                }
                                return t.next = 18, new Promise((function(t) {
                                    return setTimeout(t, r)
                                }));
                            case 18:
                                this.$fetchState.error = e, this.$fetchState.pending = !1, this.$fetchState.timestamp = Date.now(), this.$nextTick((function() {
                                    return o.$nuxt.nbFetching--
                                }));
                            case 22:
                            case "end":
                                return t.stop()
                        }
                    }), t, this, [
                        [6, 11]
                    ])
                })))).apply(this, arguments)
            }
            e.a = {
                beforeCreate: function() {
                    Object(c.l)(this) && (this._fetchDelay = "number" == typeof this.$options.fetchDelay ? this.$options.fetchDelay : 200, o.
                    default.util.defineReactive(this, "$fetchState", {
                        pending: !1,
                        error: null,
                        timestamp: Date.now()
                    }), this.$fetch = d.bind(this), Object(c.a)(this, "created", h), Object(c.a)(this, "beforeMount", f))
                }
            }
        },
        193: function(t, e, n) {
            "use strict";
            n.d(e, "i", (function() {
                return l
            })), n.d(e, "g", (function() {
                return f
            })), n.d(e, "f", (function() {
                return h
            })), n.d(e, "e", (function() {
                return d
            })), n.d(e, "c", (function() {
                return m
            })), n.d(e, "b", (function() {
                return v
            })), n.d(e, "a", (function() {
                return y
            })), n.d(e, "h", (function() {
                return w
            })), n.d(e, "d", (function() {
                return x
            }));
            n(20);
            var r = n(2),
                o = n(3),
                c = n(47),
                l = function() {
                    var t = Object(r.a)(regeneratorRuntime.mark((function t(e) {
                        var n, r, l, f, h, time, d, m, v;
                        return regeneratorRuntime.wrap((function(t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    if (n = e.keyword, r = e.page, l = e.size, f = e.ip, h = e.type, time = e.time, d = e.searchtype, t.prev = 1, null != n && null != n) {
                                        t.next = 4;
                                        break
                                    }
                                    return t.abrupt("return", null);
                                case 4:
                                    return m = {
                                        kw: n,
                                        page: r || 1,
                                        ip: f
                                    }, o.l.type[h] && (m.category = o.l.type[h].value), o.l.time[time] && (m.diffDay = o.l.time[time].value), o.l.searchtype[d] && (m.searchType = o.l.searchtype[d].value), o.l.size[l] && (o.l.size[l].minSize ? m.minSize = o.l.size[l].minSize : m.minSize = 1, o.l.size[l].maxSize ? m.maxSize = o.l.size[l].maxSize : m.maxSize = 1099511627776), t.next = 11, c.get("".concat("", "/api/search"), {
                                        params: m
                                    });
                                case 11:
                                    return v = t.sent, t.abrupt("return", v.data);
                                case 15:
                                    return t.prev = 15, t.t0 = t.
                                    catch (1), console.log(t.t0), t.abrupt("return", null);
                                case 19:
                                case "end":
                                    return t.stop()
                            }
                        }), t, null, [
                            [1, 15]
                        ])
                    })));
                    return function(e) {
                        return t.apply(this, arguments)
                    }
                }(),
                f = function() {
                    var t = Object(r.a)(regeneratorRuntime.mark((function t(e) {
                        var n, r, o, l;
                        return regeneratorRuntime.wrap((function(t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    return n = e.ip, r = e.size, o = void 0 === r ? 10 : r, t.prev = 1, t.next = 4, c.get("".concat("", "/api/query"), {
                                        params: {
                                            type: "random",
                                            size: o,
                                            ip: n
                                        }
                                    });
                                case 4:
                                    return l = t.sent, t.abrupt("return", l.data);
                                case 8:
                                    return t.prev = 8, t.t0 = t.
                                    catch (1), console.log(t.t0), t.abrupt("return", null);
                                case 12:
                                case "end":
                                    return t.stop()
                            }
                        }), t, null, [
                            [1, 8]
                        ])
                    })));
                    return function(e) {
                        return t.apply(this, arguments)
                    }
                }(),
                h = function() {
                    var t = Object(r.a)(regeneratorRuntime.mark((function t(e) {
                        var n, r, o, l;
                        return regeneratorRuntime.wrap((function(t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    return n = e.ip, r = e.size, o = void 0 === r ? 10 : r, t.prev = 1, t.next = 4, c.get("".concat("", "/api/query"), {
                                        params: {
                                            type: "new",
                                            ip: n,
                                            size: o
                                        }
                                    });
                                case 4:
                                    return l = t.sent, t.abrupt("return", l.data);
                                case 8:
                                    return t.prev = 8, t.t0 = t.
                                    catch (1), console.log(t.t0), t.abrupt("return", null);
                                case 12:
                                case "end":
                                    return t.stop()
                            }
                        }), t, null, [
                            [1, 8]
                        ])
                    })));
                    return function(e) {
                        return t.apply(this, arguments)
                    }
                }(),
                d = function() {
                    var t = Object(r.a)(regeneratorRuntime.mark((function t(e) {
                        var n, r, o, l;
                        return regeneratorRuntime.wrap((function(t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    return n = e.ip, r = e.size, o = void 0 === r ? 10 : r, t.prev = 1, t.next = 4, c.get("".concat("", "/api/query"), {
                                        params: {
                                            type: "views",
                                            ip: n,
                                            size: o
                                        }
                                    });
                                case 4:
                                    return l = t.sent, t.abrupt("return", l.data);
                                case 8:
                                    return t.prev = 8, t.t0 = t.
                                    catch (1), console.log(t.t0), t.abrupt("return", null);
                                case 12:
                                case "end":
                                    return t.stop()
                            }
                        }), t, null, [
                            [1, 8]
                        ])
                    })));
                    return function(e) {
                        return t.apply(this, arguments)
                    }
                }(),
                m = function() {
                    var t = Object(r.a)(regeneratorRuntime.mark((function t(e) {
                        var n, r, o, l;
                        return regeneratorRuntime.wrap((function(t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    return n = e.ip, r = e.size, o = void 0 === r ? 10 : r, t.prev = 1, t.next = 4, c.get("".concat("", "/api/query"), {
                                        params: {
                                            type: "download",
                                            ip: n,
                                            size: o
                                        }
                                    });
                                case 4:
                                    return l = t.sent, t.abrupt("return", l.data);
                                case 8:
                                    return t.prev = 8, t.t0 = t.
                                    catch (1), console.log(t.t0), t.abrupt("return", null);
                                case 12:
                                case "end":
                                    return t.stop()
                            }
                        }), t, null, [
                            [1, 8]
                        ])
                    })));
                    return function(e) {
                        return t.apply(this, arguments)
                    }
                }(),
                v = function() {
                    var t = Object(r.a)(regeneratorRuntime.mark((function t(e) {
                        var n, r, o, l, f;
                        return regeneratorRuntime.wrap((function(t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    return n = e.id, r = e.size, o = void 0 === r ? 15 : r, l = e.ip, t.prev = 1, t.next = 4, c.get("".concat("", "/api/detail"), {
                                        params: {
                                            id: n,
                                            size: o,
                                            ip: l
                                        }
                                    });
                                case 4:
                                    return f = t.sent, t.abrupt("return", f.data);
                                case 8:
                                    return t.prev = 8, t.t0 = t.
                                    catch (1), console.log(t.t0), t.abrupt("return", null);
                                case 12:
                                case "end":
                                    return t.stop()
                            }
                        }), t, null, [
                            [1, 8]
                        ])
                    })));
                    return function(e) {
                        return t.apply(this, arguments)
                    }
                }(),
                y = function() {
                    var t = Object(r.a)(regeneratorRuntime.mark((function t(data) {
                        var e;
                        return regeneratorRuntime.wrap((function(t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    return t.prev = 0, t.next = 3, c.post("".concat("", "/api/checkUrlValidFromBaidu"), {
                                        data: data
                                    });
                                case 3:
                                    return e = t.sent, t.abrupt("return", e.data || {});
                                case 7:
                                    return t.prev = 7, t.t0 = t.
                                    catch (0), console.log(t.t0), t.abrupt("return", {});
                                case 11:
                                case "end":
                                    return t.stop()
                            }
                        }), t, null, [
                            [0, 7]
                        ])
                    })));
                    return function(e) {
                        return t.apply(this, arguments)
                    }
                }(),
                w = function() {
                    var t = Object(r.a)(regeneratorRuntime.mark((function t() {
                        return regeneratorRuntime.wrap((function(t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    return t.abrupt("return", o.m);
                                case 1:
                                case "end":
                                    return t.stop()
                            }
                        }), t)
                    })));
                    return function() {
                        return t.apply(this, arguments)
                    }
                }(),
                x = function() {
                    var t = Object(r.a)(regeneratorRuntime.mark((function t() {
                        return regeneratorRuntime.wrap((function(t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    return t.abrupt("return", o.o);
                                case 1:
                                case "end":
                                    return t.stop()
                            }
                        }), t)
                    })));
                    return function() {
                        return t.apply(this, arguments)
                    }
                }()
        },
        194: function(t, e, n) {
            t.exports = n.p + "img/f8017ca.png"
        },
        195: function(t, e, n) {
            t.exports = n.p + "img/ac56964.png"
        },
        197: function(t, e, n) {
            t.exports = n.p + "img/a4c9c3b.png"
        },
        198: function(t, e, n) {
            t.exports = n.p + "img/a2179aa.png"
        },
        199: function(t, e, n) {
            t.exports = n(200)
        },
        200: function(t, e, n) {
            "use strict";
            n.r(e), function(t) {
                n(53), n(34), n(58), n(40), n(41), n(50);
                var e = n(48),
                    r = (n(20), n(196), n(2)),
                    o = (n(93), n(94), n(24), n(13), n(52), n(29), n(154), n(214), n(222), n(224), n(0)),
                    c = n(187),
                    l = n(128),
                    f = n(1),
                    h = n(19),
                    d = n(192),
                    m = n(92);

                function v(t, e) {
                    var n;
                    if ("undefined" == typeof Symbol || null == t[Symbol.iterator]) {
                        if (Array.isArray(t) || (n = function(t, e) {
                            if (!t) return;
                            if ("string" == typeof t) return y(t, e);
                            var n = Object.prototype.toString.call(t).slice(8, -1);
                            "Object" === n && t.constructor && (n = t.constructor.name);
                            if ("Map" === n || "Set" === n) return Array.from(t);
                            if ("Arguments" === n || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)) return y(t, e)
                        }(t)) || e && t && "number" == typeof t.length) {
                            n && (t = n);
                            var i = 0,
                                r = function() {};
                            return {
                                s: r,
                                n: function() {
                                    return i >= t.length ? {
                                        done: !0
                                    } : {
                                        done: !1,
                                        value: t[i++]
                                    }
                                },
                                e: function(t) {
                                    throw t
                                },
                                f: r
                            }
                        }
                        throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")
                    }
                    var o, c = !0,
                        l = !1;
                    return {
                        s: function() {
                            n = t[Symbol.iterator]()
                        },
                        n: function() {
                            var t = n.next();
                            return c = t.done, t
                        },
                        e: function(t) {
                            l = !0, o = t
                        },
                        f: function() {
                            try {
                                c || null == n.
                                    return ||n.
                                return ()
                            } finally {
                                if (l) throw o
                            }
                        }
                    }
                }
                function y(t, e) {
                    (null == e || e > t.length) && (e = t.length);
                    for (var i = 0, n = new Array(e); i < e; i++) n[i] = t[i];
                    return n
                }
                o.
                    default.__nuxt__fetch__mixin__ || (o.
                default.mixin(d.a), o.
                    default.__nuxt__fetch__mixin__ = !0), o.
                default.component(m.a.name, m.a), o.
                default.component("NLink", m.a), t.fetch || (t.fetch = c.a);
                var w, x, _ = [],
                    k = window.__NUXT__ || {};
                Object.assign(o.
                    default.config, {
                    silent: !0,
                    performance: !1
                });
                var j = o.
                    default.config.errorHandler || console.error;

                function O(t, e, n) {
                    for (var r = function(component) {
                        var t = function(component, t) {
                            if (!component || !component.options || !component.options[t]) return {};
                            var option = component.options[t];
                            if ("function" == typeof option) {
                                for (var e = arguments.length, n = new Array(e > 2 ? e - 2 : 0), r = 2; r < e; r++) n[r - 2] = arguments[r];
                                return option.apply(void 0, n)
                            }
                            return option
                        }(component, "transition", e, n) || {};
                        return "string" == typeof t ? {
                            name: t
                        } : t
                    }, o = n ? Object(f.g)(n) : [], c = Math.max(t.length, o.length), l = [], h = function(i) {
                        var e = Object.assign({}, r(t[i])),
                            n = Object.assign({}, r(o[i]));
                        Object.keys(e).filter((function(t) {
                            return void 0 !== e[t] && !t.toLowerCase().includes("leave")
                        })).forEach((function(t) {
                            n[t] = e[t]
                        })), l.push(n)
                    }, i = 0; i < c; i++) h(i);
                    return l
                }
                function $(t, e, n) {
                    return C.apply(this, arguments)
                }
                function C() {
                    return (C = Object(r.a)(regeneratorRuntime.mark((function t(e, n, r) {
                        var o, c, l, h, d = this;
                        return regeneratorRuntime.wrap((function(t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    if (this._routeChanged = Boolean(w.nuxt.err) || n.name !== e.name, this._paramChanged = !this._routeChanged && n.path !== e.path, this._queryChanged = !this._paramChanged && n.fullPath !== e.fullPath, this._diffQuery = this._queryChanged ? Object(f.i)(e.query, n.query) : [], (this._routeChanged || this._paramChanged) && this.$loading.start && !this.$loading.manual && this.$loading.start(), t.prev = 5, !this._queryChanged) {
                                        t.next = 12;
                                        break
                                    }
                                    return t.next = 9, Object(f.q)(e, (function(t, e) {
                                        return {
                                            Component: t,
                                            instance: e
                                        }
                                    }));
                                case 9:
                                    o = t.sent, o.some((function(t) {
                                        var r = t.Component,
                                            o = t.instance,
                                            c = r.options.watchQuery;
                                        return !0 === c || (Array.isArray(c) ? c.some((function(t) {
                                            return d._diffQuery[t]
                                        })) : "function" == typeof c && c.apply(o, [e.query, n.query]))
                                    })) && this.$loading.start && !this.$loading.manual && this.$loading.start();
                                case 12:
                                    r(), t.next = 26;
                                    break;
                                case 15:
                                    if (t.prev = 15, t.t0 = t.
                                    catch (5), c = t.t0 || {}, l = c.statusCode || c.status || c.response && c.response.status || 500, h = c.message || "", !/^Loading( CSS)? chunk (\d)+ failed\./.test(h)) {
                                        t.next = 23;
                                        break
                                    }
                                    return window.location.reload(!0), t.abrupt("return");
                                case 23:
                                    this.error({
                                        statusCode: l,
                                        message: h
                                    }), this.$nuxt.$emit("routeChanged", e, n, c), r();
                                case 26:
                                case "end":
                                    return t.stop()
                            }
                        }), t, this, [
                            [5, 15]
                        ])
                    })))).apply(this, arguments)
                }
                function S(t, e) {
                    return k.serverRendered && e && Object(f.b)(t, e), t._Ctor = t, t
                }
                function R(t) {
                    var path = Object(f.f)(t.options.base, t.options.mode);
                    return Object(f.d)(t.match(path), function() {
                        var t = Object(r.a)(regeneratorRuntime.mark((function t(e, n, r, o, c) {
                            var l;
                            return regeneratorRuntime.wrap((function(t) {
                                for (;;) switch (t.prev = t.next) {
                                    case 0:
                                        if ("function" != typeof e || e.options) {
                                            t.next = 4;
                                            break
                                        }
                                        return t.next = 3, e();
                                    case 3:
                                        e = t.sent;
                                    case 4:
                                        return l = S(Object(f.r)(e), k.data ? k.data[c] : null), r.components[o] = l, t.abrupt("return", l);
                                    case 7:
                                    case "end":
                                        return t.stop()
                                }
                            }), t)
                        })));
                        return function(e, n, r, o, c) {
                            return t.apply(this, arguments)
                        }
                    }())
                }
                function E(t, e, n) {
                    var r = this,
                        o = [],
                        c = !1;
                    if (void 0 !== n && (o = [], (n = Object(f.r)(n)).options.middleware && (o = o.concat(n.options.middleware)), t.forEach((function(t) {
                        t.options.middleware && (o = o.concat(t.options.middleware))
                    }))), o = o.map((function(t) {
                        return "function" == typeof t ? t : ("function" != typeof l.a[t] && (c = !0, r.error({
                            statusCode: 500,
                            message: "Unknown middleware " + t
                        })), l.a[t])
                    })), !c) return Object(f.n)(o, e)
                }
                function T(t, e, n) {
                    return A.apply(this, arguments)
                }
                function A() {
                    return (A = Object(r.a)(regeneratorRuntime.mark((function t(e, n, r) {
                        var o, c, l, d, m, y, x, k, j, $, C, S, R, T, A, P = this;
                        return regeneratorRuntime.wrap((function(t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    if (!1 !== this._routeChanged || !1 !== this._paramChanged || !1 !== this._queryChanged) {
                                        t.next = 2;
                                        break
                                    }
                                    return t.abrupt("return", r());
                                case 2:
                                    return e === n ? _ = [] : (o = [], _ = Object(f.g)(n, o).map((function(t, i) {
                                        return Object(f.c)(n.matched[o[i]].path)(n.params)
                                    }))), c = !1, l = function(path) {
                                        n.path === path.path && P.$loading.finish && P.$loading.finish(), n.path !== path.path && P.$loading.pause && P.$loading.pause(), c || (c = !0, r(path))
                                    }, t.next = 7, Object(f.s)(w, {
                                        route: e,
                                        from: n,
                                        next: l.bind(this)
                                    });
                                case 7:
                                    if (this._dateLastError = w.nuxt.dateErr, this._hadError = Boolean(w.nuxt.err), d = [], (m = Object(f.g)(e, d)).length) {
                                        t.next = 26;
                                        break
                                    }
                                    return t.next = 14, E.call(this, m, w.context);
                                case 14:
                                    if (!c) {
                                        t.next = 16;
                                        break
                                    }
                                    return t.abrupt("return");
                                case 16:
                                    return y = (h.a.options || h.a).layout, t.next = 19, this.loadLayout("function" == typeof y ? y.call(h.a, w.context) : y);
                                case 19:
                                    return x = t.sent, t.next = 22, E.call(this, m, w.context, x);
                                case 22:
                                    if (!c) {
                                        t.next = 24;
                                        break
                                    }
                                    return t.abrupt("return");
                                case 24:
                                    return w.context.error({
                                        statusCode: 404,
                                        message: "This page could not be found"
                                    }), t.abrupt("return", r());
                                case 26:
                                    return m.forEach((function(t) {
                                        t._Ctor && t._Ctor.options && (t.options.asyncData = t._Ctor.options.asyncData, t.options.fetch = t._Ctor.options.fetch)
                                    })), this.setTransitions(O(m, e, n)), t.prev = 28, t.next = 31, E.call(this, m, w.context);
                                case 31:
                                    if (!c) {
                                        t.next = 33;
                                        break
                                    }
                                    return t.abrupt("return");
                                case 33:
                                    if (!w.context._errored) {
                                        t.next = 35;
                                        break
                                    }
                                    return t.abrupt("return", r());
                                case 35:
                                    return "function" == typeof(k = m[0].options.layout) && (k = k(w.context)), t.next = 39, this.loadLayout(k);
                                case 39:
                                    return k = t.sent, t.next = 42, E.call(this, m, w.context, k);
                                case 42:
                                    if (!c) {
                                        t.next = 44;
                                        break
                                    }
                                    return t.abrupt("return");
                                case 44:
                                    if (!w.context._errored) {
                                        t.next = 46;
                                        break
                                    }
                                    return t.abrupt("return", r());
                                case 46:
                                    j = !0, t.prev = 47, $ = v(m), t.prev = 49, $.s();
                                case 51:
                                    if ((C = $.n()).done) {
                                        t.next = 62;
                                        break
                                    }
                                    if ("function" == typeof(S = C.value).options.validate) {
                                        t.next = 55;
                                        break
                                    }
                                    return t.abrupt("continue", 60);
                                case 55:
                                    return t.next = 57, S.options.validate(w.context);
                                case 57:
                                    if (j = t.sent) {
                                        t.next = 60;
                                        break
                                    }
                                    return t.abrupt("break", 62);
                                case 60:
                                    t.next = 51;
                                    break;
                                case 62:
                                    t.next = 67;
                                    break;
                                case 64:
                                    t.prev = 64, t.t0 = t.
                                    catch (49), $.e(t.t0);
                                case 67:
                                    return t.prev = 67, $.f(), t.finish(67);
                                case 70:
                                    t.next = 76;
                                    break;
                                case 72:
                                    return t.prev = 72, t.t1 = t.
                                    catch (47), this.error({
                                        statusCode: t.t1.statusCode || "500",
                                        message: t.t1.message
                                    }), t.abrupt("return", r());
                                case 76:
                                    if (j) {
                                        t.next = 79;
                                        break
                                    }
                                    return this.error({
                                        statusCode: 404,
                                        message: "This page could not be found"
                                    }), t.abrupt("return", r());
                                case 79:
                                    return t.next = 81, Promise.all(m.map((function(t, i) {
                                        t._path = Object(f.c)(e.matched[d[i]].path)(e.params), t._dataRefresh = !1;
                                        var r = t._path !== _[i];
                                        if (P._routeChanged && r) t._dataRefresh = !0;
                                        else if (P._paramChanged && r) {
                                            var o = t.options.watchParam;
                                            t._dataRefresh = !1 !== o
                                        } else if (P._queryChanged) {
                                            var c = t.options.watchQuery;
                                            !0 === c ? t._dataRefresh = !0 : Array.isArray(c) ? t._dataRefresh = c.some((function(t) {
                                                return P._diffQuery[t]
                                            })) : "function" == typeof c && (R || (R = Object(f.h)(e)), t._dataRefresh = c.apply(R[i], [e.query, n.query]))
                                        }
                                        if (P._hadError || !P._isMounted || t._dataRefresh) {
                                            var l = [],
                                                h = t.options.asyncData && "function" == typeof t.options.asyncData,
                                                m = Boolean(t.options.fetch) && t.options.fetch.length,
                                                v = h && m ? 30 : 45;
                                            if (h) {
                                                var y = Object(f.p)(t.options.asyncData, w.context).then((function(e) {
                                                    Object(f.b)(t, e), P.$loading.increase && P.$loading.increase(v)
                                                }));
                                                l.push(y)
                                            }
                                            if (P.$loading.manual = !1 === t.options.loading, m) {
                                                var p = t.options.fetch(w.context);
                                                p && (p instanceof Promise || "function" == typeof p.then) || (p = Promise.resolve(p)), p.then((function(t) {
                                                    P.$loading.increase && P.$loading.increase(v)
                                                })), l.push(p)
                                            }
                                            return Promise.all(l)
                                        }
                                    })));
                                case 81:
                                    c || (this.$loading.finish && !this.$loading.manual && this.$loading.finish(), r()), t.next = 98;
                                    break;
                                case 84:
                                    if (t.prev = 84, t.t2 = t.
                                    catch (28), "ERR_REDIRECT" !== (T = t.t2 || {}).message) {
                                        t.next = 89;
                                        break
                                    }
                                    return t.abrupt("return", this.$nuxt.$emit("routeChanged", e, n, T));
                                case 89:
                                    return _ = [], Object(f.k)(T), "function" == typeof(A = (h.a.options || h.a).layout) && (A = A(w.context)), t.next = 95, this.loadLayout(A);
                                case 95:
                                    this.error(T), this.$nuxt.$emit("routeChanged", e, n, T), r();
                                case 98:
                                case "end":
                                    return t.stop()
                            }
                        }), t, this, [
                            [28, 84],
                            [47, 72],
                            [49, 64, 67, 70]
                        ])
                    })))).apply(this, arguments)
                }
                function P(t, n) {
                    Object(f.d)(t, (function(t, n, r, c) {
                        return "object" !== Object(e.a)(t) || t.options || ((t = o.
                        default.extend(t))._Ctor = t, r.components[c] = t), t
                    }))
                }
                function N(t) {
                    this._hadError && this._dateLastError === this.$options.nuxt.dateErr && this.error();
                    var e = this.$options.nuxt.err ? (h.a.options || h.a).layout : t.matched[0].components.
                        default.options.layout;
                    "function" == typeof e && (e = e(w.context)), this.setLayout(e)
                }
                function z(t, e) {
                    var n = this;
                    if (!1 !== this._routeChanged || !1 !== this._paramChanged || !1 !== this._queryChanged) {
                        var r = Object(f.h)(t),
                            c = Object(f.g)(t);
                        o.
                        default.nextTick((function() {
                            r.forEach((function(t, i) {
                                if (t && !t._isDestroyed && t.constructor._dataRefresh && c[i] === t.constructor && !0 !== t.$vnode.data.keepAlive && "function" == typeof t.constructor.options.data) {
                                    var e = t.constructor.options.data.call(t);
                                    for (var n in e) o.
                                    default.set(t.$data, n, e[n]);
                                    window.$nuxt.$nextTick((function() {
                                        window.$nuxt.$emit("triggerScroll")
                                    }))
                                }
                            })), N.call(n, t)
                        }))
                    }
                }
                function D(t) {
                    window.onNuxtReadyCbs.forEach((function(e) {
                        "function" == typeof e && e(t)
                    })), "function" == typeof window._onNuxtLoaded && window._onNuxtLoaded(t), x.afterEach((function(e, n) {
                        o.
                        default.nextTick((function() {
                            return t.$nuxt.$emit("routeChanged", e, n)
                        }))
                    }))
                }
                function L() {
                    return (L = Object(r.a)(regeneratorRuntime.mark((function t(e) {
                        var n, r, c, l, h;
                        return regeneratorRuntime.wrap((function(t) {
                            for (;;) switch (t.prev = t.next) {
                                case 0:
                                    return w = e.app, x = e.router, e.store, n = new o.
                                    default (w), r = k.layout || "default", t.next = 7, n.loadLayout(r);
                                case 7:
                                    return n.setLayout(r), c = function() {
                                        n.$mount("#__nuxt"), x.afterEach(P), x.afterEach(z.bind(n)), o.
                                        default.nextTick((function() {
                                            D(n)
                                        }))
                                    }, t.next = 11, Promise.all(R(x));
                                case 11:
                                    if (l = t.sent, n.setTransitions = n.$options.nuxt.setTransitions.bind(n), l.length && (n.setTransitions(O(l, x.currentRoute)), _ = x.currentRoute.matched.map((function(t) {
                                        return Object(f.c)(t.path)(x.currentRoute.params)
                                    }))), n.$loading = {}, k.error && n.error(k.error), x.beforeEach($.bind(n)), x.beforeEach(T.bind(n)), !k.serverRendered || k.routePath !== n.context.route.path) {
                                        t.next = 21;
                                        break
                                    }
                                    return c(), t.abrupt("return");
                                case 21:
                                    h = function() {
                                        P(x.currentRoute, x.currentRoute), N.call(n, x.currentRoute), c()
                                    }, T.call(n, x.currentRoute, x.currentRoute, (function(path) {
                                        if (path) {
                                            var t = x.afterEach((function(e, n) {
                                                t(), h()
                                            }));
                                            x.push(path, void 0, (function(t) {
                                                t && j(t)
                                            }))
                                        } else h()
                                    }));
                                case 23:
                                case "end":
                                    return t.stop()
                            }
                        }), t)
                    })))).apply(this, arguments)
                }
                Object(h.b)().then((function(t) {
                    return L.apply(this, arguments)
                })).
                catch (j)
            }.call(this, n(25))
        },
        236: function(t, e, n) {
            "use strict";
            var r = n(85);
            n.n(r).a
        },
        237: function(t, e, n) {
            "use strict";
            var r = n(86);
            n.n(r).a
        },
        238: function(t, e, n) {
            "use strict";
            var r = n(87);
            n.n(r).a
        },
        255: function(t, e, n) {
            "use strict";
            var r = n(88);
            n.n(r).a
        },
        256: function(t, e, n) {
            "use strict";
            var r = n(89);
            n.n(r).a
        },
        258: function(t, e, n) {
            t.exports = n.p + "img/5fe16ff.jpg"
        },
        259: function(t, e, n) {
            "use strict";
            var r = n(90);
            n.n(r).a
        },
        260: function(t, e, n) {
            "use strict";
            var r = n(91);
            n.n(r).a
        },
        261: function(t, e, n) {
            "use strict";
            n.r(e), n.d(e, "state", (function() {
                return r
            })), n.d(e, "mutations", (function() {
                return o
            }));
            var r = function() {
                    return {
                        filter: {
                            show: !1,
                            time: -1,
                            type: -1,
                            size: -1,
                            searchtype: -1
                        }
                    }
                },
                o = {
                    changeFilter: function(t, filter) {
                        t.filter = filter
                    },
                    showFilter: function(t) {
                        t.filter.show = !0
                    },
                    hideFilter: function(t) {
                        t.filter.show = !1
                    },
                    changeFilterSize: function(t, e) {
                        null != e && (t.filter.size = e)
                    },
                    changeFilterTime: function(t, e) {
                        null != e && (t.filter.time = e)
                    },
                    changeFilterType: function(t, e) {
                        null != e && (t.filter.type = e)
                    },
                    changeFilterSearchType: function(t, e) {
                        null != e && (t.filter.searchtype = e)
                    }
                }
        },
        3: function(t, e, n) {
            "use strict";
            n.d(e, "i", (function() {
                return r
            })), n.d(e, "c", (function() {
                return o
            })), n.d(e, "g", (function() {
                return c
            })), n.d(e, "j", (function() {
                return l
            })), n.d(e, "d", (function() {
                return f
            })), n.d(e, "f", (function() {
                return h
            })), n.d(e, "k", (function() {
                return d
            })), n.d(e, "e", (function() {
                return m
            })), n.d(e, "h", (function() {
                return v
            })), n.d(e, "p", (function() {
                return y
            })), n.d(e, "b", (function() {
                return w
            })), n.d(e, "r", (function() {
                return x
            })), n.d(e, "n", (function() {
                return _
            })), n.d(e, "a", (function() {
                return k
            })), n.d(e, "q", (function() {
                return j
            })), n.d(e, "l", (function() {
                return O
            })), n.d(e, "o", (function() {
                return $
            })), n.d(e, "m", (function() {
                return C
            }));
            var r = "filetype_video",
                o = "filetype_audio",
                c = "filetype_ppt",
                l = "filetype_xls",
                f = "filetype_doc",
                h = "filetype_pdf",
                d = "filetype_zip",
                m = "filetype_img",
                v = "filetype_unknown",
                y = "大力盘搜索",
                w = "https://www.dalipan.com",
                x = "网盘搜索，就用".concat(y, " - 最好用的百度网盘搜索引擎，").concat(w, "。"),
                _ = "".concat(y, ",盘搜搜,大力盘,大力盘搜索,大力搜索盘,网盘搜索,电影下载,迅雷下载,bt下载,种子下载,电子书下载,百度云盘搜索,网盘搜索引擎,百度网盘搜索"),
                k = "".concat(y, "支持百度云搜索，可快速搜索百度网盘资源中的有效连接，自动识别无效的百度云网盘资源，每天更新海量资源。"),
                j = {
                    sopan_spider: {
                        name: "52搜盘",
                        url: "http://www.52sopan.com"
                    },
                    panduoduo: {
                        name: "盘多多",
                        url: "http://www.panduoduo.net"
                    },
                    "56wangpan": {
                        name: "56网盘",
                        url: "http://www.56wanpan.com"
                    },
                    pansoso: {
                        name: "盘搜搜",
                        url: "http://www.pansoso.com"
                    },
                    xiaobaipan: {
                        name: "小白盘",
                        url: "http://xiaobaipan.com"
                    },
                    quzhuanpan: {
                        name: "去转盘",
                        url: "https://www.quzhuanpan.com"
                    }
                },
                O = {
                    type: {
                        7: {
                            name: "BT种子",
                            value: 7
                        },
                        6: {
                            name: "压缩包",
                            value: 6
                        },
                        5: {
                            name: "软件",
                            value: 5
                        },
                        4: {
                            name: "文档",
                            value: 4
                        },
                        3: {
                            name: "图片",
                            value: 3
                        },
                        2: {
                            name: "音乐",
                            value: 2
                        },
                        1: {
                            name: "视频",
                            value: 1
                        },
                        0: {
                            name: "文件夹",
                            value: 0
                        }
                    },
                    time: {
                        3: {
                            name: "最近一年",
                            value: 365
                        },
                        2: {
                            name: "最近半年",
                            value: 180
                        },
                        1: {
                            name: "最近一月",
                            value: 30
                        },
                        0: {
                            name: "最近一周",
                            value: 7
                        }
                    },
                    size: {
                        3: {
                            name: "大于2GB",
                            minSize: 2147483648
                        },
                        2: {
                            name: "200MB-2GB",
                            maxSize: 2147483648,
                            minSize: 209715200
                        },
                        1: {
                            name: "20MB-200MB",
                            maxSize: 209715200,
                            minSize: 20971520
                        },
                        0: {
                            name: "小于20MB",
                            maxSize: 20971520
                        }
                    },
                    searchtype: {
                        1: {
                            name: "模糊搜索",
                            value: "match"
                        },
                        0: {
                            name: "精确搜索",
                            value: "precise"
                        }
                    }
                },
                $ = [{
                    text: "北邮人导航",
                    link: "http://byr.wiki?from=dalipan"
                }, {
                    text: "柴杜导航",
                    link: "https://www.chaidu.com?from=dalipan"
                }, {
                    text: "24K导航",
                    link: "https://www.24kdh.com?from=dalipan"
                }],
                C = ["课程", "资料", "考研", "学习", "书单", "英语", "小说", "漫画", "电影", "游戏", "电子书", "美女", "python"]
        },
        85: function(t, e, n) {},
        86: function(t, e, n) {},
        87: function(t, e, n) {},
        88: function(t, e, n) {},
        89: function(t, e, n) {},
        90: function(t, e, n) {},
        91: function(t, e, n) {},
        92: function(t, e, n) {
            "use strict";
            n(24), n(53), n(34), n(50), n(58), n(29), n(40), n(41), n(13), n(93), n(94);
            var r = n(0);

            function o(t, e) {
                var n;
                if ("undefined" == typeof Symbol || null == t[Symbol.iterator]) {
                    if (Array.isArray(t) || (n = function(t, e) {
                        if (!t) return;
                        if ("string" == typeof t) return c(t, e);
                        var n = Object.prototype.toString.call(t).slice(8, -1);
                        "Object" === n && t.constructor && (n = t.constructor.name);
                        if ("Map" === n || "Set" === n) return Array.from(t);
                        if ("Arguments" === n || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)) return c(t, e)
                    }(t)) || e && t && "number" == typeof t.length) {
                        n && (t = n);
                        var i = 0,
                            r = function() {};
                        return {
                            s: r,
                            n: function() {
                                return i >= t.length ? {
                                    done: !0
                                } : {
                                    done: !1,
                                    value: t[i++]
                                }
                            },
                            e: function(t) {
                                throw t
                            },
                            f: r
                        }
                    }
                    throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")
                }
                var o, l = !0,
                    f = !1;
                return {
                    s: function() {
                        n = t[Symbol.iterator]()
                    },
                    n: function() {
                        var t = n.next();
                        return l = t.done, t
                    },
                    e: function(t) {
                        f = !0, o = t
                    },
                    f: function() {
                        try {
                            l || null == n.
                                return ||n.
                            return ()
                        } finally {
                            if (f) throw o
                        }
                    }
                }
            }
            function c(t, e) {
                (null == e || e > t.length) && (e = t.length);
                for (var i = 0, n = new Array(e); i < e; i++) n[i] = t[i];
                return n
            }
            var l = window.requestIdleCallback ||
                function(t) {
                    var e = Date.now();
                    return setTimeout((function() {
                        t({
                            didTimeout: !1,
                            timeRemaining: function() {
                                return Math.max(0, 50 - (Date.now() - e))
                            }
                        })
                    }), 1)
                }, f = window.cancelIdleCallback ||
                function(t) {
                    clearTimeout(t)
                }, h = window.IntersectionObserver && new window.IntersectionObserver((function(t) {
                t.forEach((function(t) {
                    var e = t.intersectionRatio,
                        link = t.target;
                    e <= 0 || link.__prefetch()
                }))
            }));
            e.a = {
                name: "NuxtLink",
                extends: r.
                default.component("RouterLink"),
                props:
                    {
                        prefetch: {
                            type: Boolean,
                            default:
                                !0
                        },
                        noPrefetch: {
                            type: Boolean,
                            default:
                                !1
                        }
                    },
                mounted: function() {
                    this.prefetch && !this.noPrefetch && (this.handleId = l(this.observe, {
                        timeout: 2e3
                    }))
                },
                beforeDestroy: function() {
                    f(this.handleId), this.__observed && (h.unobserve(this.$el), delete this.$el.__prefetch)
                },
                methods: {
                    observe: function() {
                        h && this.shouldPrefetch() && (this.$el.__prefetch = this.prefetchLink.bind(this), h.observe(this.$el), this.__observed = !0)
                    },
                    shouldPrefetch: function() {
                        return this.getPrefetchComponents().length > 0
                    },
                    canPrefetch: function() {
                        var t = navigator.connection;
                        return !(this.$nuxt.isOffline || t && ((t.effectiveType || "").includes("2g") || t.saveData))
                    },
                    getPrefetchComponents: function() {
                        return this.$router.resolve(this.to, this.$route, this.append).resolved.matched.map((function(t) {
                            return t.components.
                                default
                        })).filter((function(t) {
                            return "function" == typeof t && !t.options && !t.__prefetched
                        }))
                    },
                    prefetchLink: function() {
                        if (this.canPrefetch()) {
                            h.unobserve(this.$el);
                            var t, e = o(this.getPrefetchComponents());
                            try {
                                for (e.s(); !(t = e.n()).done;) {
                                    var n = t.value,
                                        r = n();
                                    r instanceof Promise && r.
                                    catch ((function() {})), n.__prefetched = !0
                                }
                            } catch (t) {
                                e.e(t)
                            } finally {
                                e.f()
                            }
                        }
                    }
                }
            }
        }
    }, [
        [199, 11, 12]
    ]
]);