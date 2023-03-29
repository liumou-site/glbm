# -*- encoding: utf-8 -*-
import os

res = os.system("go test ./")
if int(res) == 0:
    print("测试通过")
else:
    print("测试失败")
