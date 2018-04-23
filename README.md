# esa_cli

Last update: Mon Apr 23 07:42:51 UTC 2018

## description

Command line tools for [esa.io](https://esa.io/).
Based on [go-esa](https://github.com/upamune/go-esa) library.

`esa_get` and `esa_new` commands create following files.

~~~bash
$ esa_new "path/to/category" "post title"
$ ls -F
$ 352/
$ tree 352/
352/
├── BodyMd
├── Category
├── Message
├── Name
├── Tags
└── Wip
~~~

You can edit these files in your local.

`esa_update` posts local updates to server.

~~~
$ vi 352/BodyMd
(edit...)
$ esa_update 352
~~~

## configuration

set these environment variables:

- `ESA_API_KEY`
  - application token
- `ESA_TEAM`
  - your team
- `ESA_USER`
  - your user name

## usage

- `esa_new` CATEGORY NAME
  - create new post
  - args
    - CATEGORY: /path/to/category/
    - NAME: post title
- `esa_update` NUMBER
  - update post
  - args
    - NUMBER: post number
- `esa_get` NUMBER
  - get specified posts
  - args
    - NUMBER: post number
- `esa_getall`
  - get all posts related to your user name

## license
The MIT License (MIT)

Copyright (c) 2018 haru (uobikiemukot at gmail dot com)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
