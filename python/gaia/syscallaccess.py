# -*- coding: utf-8 -*-

from pyelemental import RESTObject

class SyscallAccess(RESTObject):
    """ Represents a SyscallAccess in the 

        Notes:
            None
    """

    def __init__(self, **kwargs):
        """ Initializes a SyscallAccess instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> syscallaccess = SyscallAccess(id=u'xxxx-xxx-xxx-xxx', name=u'SyscallAccess')
              >>> syscallaccess = SyscallAccess(data=my_dict)
        """

        super(SyscallAccess, self).__init__()

        # Read/Write Attributes
        
        self._pid = None
        self._propagate = None
        self._action = None
        self._count = None
        self._name = None
        self._processname = None
        
        self.expose_attribute(local_name="PID", remote_name="PID")
        self.expose_attribute(local_name="Propagate", remote_name="Propagate")
        self.expose_attribute(local_name="action", remote_name="action")
        self.expose_attribute(local_name="count", remote_name="count")
        self.expose_attribute(local_name="name", remote_name="name")
        self.expose_attribute(local_name="processName", remote_name="processName")

        self._compute_args(**kwargs)

    def __str__(self):
        return '<%s:%s>' % (self.identity()["name"], self.identifier)

    def identifier(self):
        """ Identifier returns the value of the object's unique identifier.
        """
        return ""
        

    def setIdentifier(self, ID):
        """ SetIdentifier sets the value of the object's unique identifier.
        """
        pass
        

    def identity(self):
        """ Identity returns the Identity of the object.
        """
        return syscallaccessIdentity

    # Properties
    @property
    def PID(self):
        """ Get PID value.

          Notes:
              PID is the PID of the process that used the system call.

              
        """
        return self._pid

    @PID.setter
    def PID(self, value):
        """ Set PID value.

          Notes:
              PID is the PID of the process that used the system call.

              
        """
        self._pid = value
    
    @property
    def Propagate(self):
        """ Get Propagate value.

          Notes:
              Propagate indicates that the policy must be propagated to the children namespaces.

              
        """
        return self._propagate

    @Propagate.setter
    def Propagate(self, value):
        """ Set Propagate value.

          Notes:
              Propagate indicates that the policy must be propagated to the children namespaces.

              
        """
        self._propagate = value
    
    @property
    def action(self):
        """ Get action value.

          Notes:
              Actions tells if the system call has been allowed.

              
        """
        return self._action

    @action.setter
    def action(self, value):
        """ Set action value.

          Notes:
              Actions tells if the system call has been allowed.

              
        """
        self._action = value
    
    @property
    def count(self):
        """ Get count value.

          Notes:
              Count tells how many times the syscall has been sent.

              
        """
        return self._count

    @count.setter
    def count(self, value):
        """ Set count value.

          Notes:
              Count tells how many times the syscall has been sent.

              
        """
        self._count = value
    
    @property
    def name(self):
        """ Get name value.

          Notes:
              Name represents the name of the system call.

              
        """
        return self._name

    @name.setter
    def name(self, value):
        """ Set name value.

          Notes:
              Name represents the name of the system call.

              
        """
        self._name = value
    
    @property
    def processName(self):
        """ Get processName value.

          Notes:
              ProcessName is the name of the process that used the system call.

              
        """
        return self._processname

    @processName.setter
    def processName(self, value):
        """ Set processName value.

          Notes:
              ProcessName is the name of the process that used the system call.

              
        """
        self._processname = value
    

    # syscallaccessIdentity represents the Identity of the object
syscallaccessIdentity = {"name": "syscallaccess", "category": "syscallaccesses", "constructor": SyscallAccess}