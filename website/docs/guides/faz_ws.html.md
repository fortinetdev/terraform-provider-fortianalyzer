---
subcategory: ""
layout: "fortianalyzer"
page_title: "To Set the Permission Level for RPC-Permit"
description: |-
  To Set the Permission Level for RPC-Permit
---


# To Set the Permission Level for RPC-Permit

Set the permission level for RPC-Permit:

```
config system admin user
    edit "apiusertest"
        set password "your_password"
        set rpc-permit read-write
        set profileid Super_User
    end
end
```

FortiAnalyzer is now configured and ready to accept terraform requests.

The user name is `apiusertest`, and password is `your_password` in this example. **Please specify your own password!**