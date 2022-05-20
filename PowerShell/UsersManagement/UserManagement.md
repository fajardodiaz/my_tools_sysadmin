# User Management in Windows Server

## Get local users
C:> get-localuser


## Get local groups
C:> get-localgroup


## Get the members of a specific groups
C:> get-localgroupmember 'Group Name'


## Create a new local user using the New-LocalUser cmdlet
    ### First, you can create secure password variable
        C:> $Password = Read-Host -AsSecureString
        C:> New-LocalUser -Name 'user name' -Description "A simple description" -Password $Password


## Resseting a user Password
    ### First, you can create secure password variable
        C:> $Password = Read-Host -AsSecureString
    
    ### Second, you can assign the properties of the user account that you need to reset the password 
        C:> $UserAccount = Get-LocalUser -Name 'User Name'

    ### Finally, you can write the next command to reset the password
        C:> $UserAccount | Set-LocalUser -Password $Password



## Add a User to a localgroup
    C:> Add-LocalGroupMember -Group "Group Name" -Member "User Name"


## Remove a User or group from a local group
    C:> Remove-LocalGroupMember -Group "Group Name" -Member "User Name"
