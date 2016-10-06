# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class ClairNotification(RESTObject):
    """ Represents a ClairNotification in the 

        Notes:
            None
    """

    def __init__(self, **kwargs):
        """ Initializes a ClairNotification instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> clairnotification = ClairNotification(id=u'xxxx-xxx-xxx-xxx', name=u'ClairNotification')
              >>> clairnotification = ClairNotification(data=my_dict)
        """

        super(ClairNotification, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._annotation = None
        self._associatedtags = None
        self._created = None
        self._createdat = None
        self._deleted = None
        self._layerlimit = None
        self._layersintroducingnewvulnerability = None
        self._layersintroducingoldvulnerability = None
        self._name = None
        self._namespace = None
        self._newvulnerabilitylink = None
        self._newvulnerabilityname = None
        self._nextpage = None
        self._notification = None
        self._notified = None
        self._oldvulnerabilitylink = None
        self._oldvulnerabilityname = None
        self._page = None
        self._parentid = None
        self._parenttype = None
        self._status = None
        self._updatedat = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="annotation", remote_name="annotation")
        self.expose_attribute(local_name="associatedTags", remote_name="associatedTags")
        self.expose_attribute(local_name="created", remote_name="created")
        self.expose_attribute(local_name="createdAt", remote_name="createdAt")
        self.expose_attribute(local_name="deleted", remote_name="deleted")
        self.expose_attribute(local_name="layerlimit", remote_name="layerlimit")
        self.expose_attribute(local_name="layersIntroducingNewVulnerability", remote_name="layersIntroducingNewVulnerability")
        self.expose_attribute(local_name="layersIntroducingOldVulnerability", remote_name="layersIntroducingOldVulnerability")
        self.expose_attribute(local_name="name", remote_name="name")
        self.expose_attribute(local_name="namespace", remote_name="namespace")
        self.expose_attribute(local_name="newVulnerabilityLink", remote_name="newVulnerabilityLink")
        self.expose_attribute(local_name="newVulnerabilityName", remote_name="newVulnerabilityName")
        self.expose_attribute(local_name="nextPage", remote_name="nextPage")
        self.expose_attribute(local_name="notification", remote_name="notification")
        self.expose_attribute(local_name="notified", remote_name="notified")
        self.expose_attribute(local_name="oldVulnerabilityLink", remote_name="oldVulnerabilityLink")
        self.expose_attribute(local_name="oldVulnerabilityName", remote_name="oldVulnerabilityName")
        self.expose_attribute(local_name="page", remote_name="page")
        self.expose_attribute(local_name="parentID", remote_name="parentID")
        self.expose_attribute(local_name="parentType", remote_name="parentType")
        self.expose_attribute(local_name="status", remote_name="status")
        self.expose_attribute(local_name="updatedAt", remote_name="updatedAt")

        self._compute_args(**kwargs)

    def __str__(self):
        return '<%s:%s>' % (self.identity()["name"], self.identifier)

    def identifier(self):
        """ Identifier returns the value of the object's unique identifier.
        """
        
        return self.ID
        

    def setIdentifier(self, ID):
        """ SetIdentifier sets the value of the object's unique identifier.
        """
        
        self.ID = ID
        

    def identity(self):
        """ Identity returns the Identity of the object.
        """
        return clairnotificationIdentity

    # Properties
    @property
    def ID(self):
        """ Get ID value.

          Notes:
              ID is the identifier of the object.

              
        """
        return self._id

    @ID.setter
    def ID(self, value):
        """ Set ID value.

          Notes:
              ID is the identifier of the object.

              
        """
        self._id = value
    
    @property
    def annotation(self):
        """ Get annotation value.

          Notes:
              Annotation stores additional information about an entity

              
        """
        return self._annotation

    @annotation.setter
    def annotation(self, value):
        """ Set annotation value.

          Notes:
              Annotation stores additional information about an entity

              
        """
        self._annotation = value
    
    @property
    def associatedTags(self):
        """ Get associatedTags value.

          Notes:
              AssociatedTags are the list of tags attached to an entity

              
        """
        return self._associatedtags

    @associatedTags.setter
    def associatedTags(self, value):
        """ Set associatedTags value.

          Notes:
              AssociatedTags are the list of tags attached to an entity

              
        """
        self._associatedtags = value
    
    @property
    def created(self):
        """ Get created value.

          Notes:
              Created is the time when then notification was created

              
        """
        return self._created

    @created.setter
    def created(self, value):
        """ Set created value.

          Notes:
              Created is the time when then notification was created

              
        """
        self._created = value
    
    @property
    def createdAt(self):
        """ Get createdAt value.

          Notes:
              CreatedAt is the time at which an entity was created

              
        """
        return self._createdat

    @createdAt.setter
    def createdAt(self, value):
        """ Set createdAt value.

          Notes:
              CreatedAt is the time at which an entity was created

              
        """
        self._createdat = value
    
    @property
    def deleted(self):
        """ Get deleted value.

          Notes:
              Deleted is the time when the notification was deleted

              
        """
        return self._deleted

    @deleted.setter
    def deleted(self, value):
        """ Set deleted value.

          Notes:
              Deleted is the time when the notification was deleted

              
        """
        self._deleted = value
    
    @property
    def layerlimit(self):
        """ Get layerlimit value.

          Notes:
              LayerLimit is the number of layers returned in notification

              
        """
        return self._layerlimit

    @layerlimit.setter
    def layerlimit(self, value):
        """ Set layerlimit value.

          Notes:
              LayerLimit is the number of layers returned in notification

              
        """
        self._layerlimit = value
    
    @property
    def layersIntroducingNewVulnerability(self):
        """ Get layersIntroducingNewVulnerability value.

          Notes:
              LayersIntroducingNewVulnerability defines layers that are effected by new vulnerability

              
        """
        return self._layersintroducingnewvulnerability

    @layersIntroducingNewVulnerability.setter
    def layersIntroducingNewVulnerability(self, value):
        """ Set layersIntroducingNewVulnerability value.

          Notes:
              LayersIntroducingNewVulnerability defines layers that are effected by new vulnerability

              
        """
        self._layersintroducingnewvulnerability = value
    
    @property
    def layersIntroducingOldVulnerability(self):
        """ Get layersIntroducingOldVulnerability value.

          Notes:
              LayersIntroducingOldVulnerability defines layers that are effected by old vulnerability

              
        """
        return self._layersintroducingoldvulnerability

    @layersIntroducingOldVulnerability.setter
    def layersIntroducingOldVulnerability(self, value):
        """ Set layersIntroducingOldVulnerability value.

          Notes:
              LayersIntroducingOldVulnerability defines layers that are effected by old vulnerability

              
        """
        self._layersintroducingoldvulnerability = value
    
    @property
    def name(self):
        """ Get name value.

          Notes:
              Name is the name of the notification

              
        """
        return self._name

    @name.setter
    def name(self, value):
        """ Set name value.

          Notes:
              Name is the name of the notification

              
        """
        self._name = value
    
    @property
    def namespace(self):
        """ Get namespace value.

          Notes:
              Namespace tag attached to an entity

              
        """
        return self._namespace

    @namespace.setter
    def namespace(self, value):
        """ Set namespace value.

          Notes:
              Namespace tag attached to an entity

              
        """
        self._namespace = value
    
    @property
    def newVulnerabilityLink(self):
        """ Get newVulnerabilityLink value.

          Notes:
              NewVulnerabilityLink is the link that point to the new vulnerability

              
        """
        return self._newvulnerabilitylink

    @newVulnerabilityLink.setter
    def newVulnerabilityLink(self, value):
        """ Set newVulnerabilityLink value.

          Notes:
              NewVulnerabilityLink is the link that point to the new vulnerability

              
        """
        self._newvulnerabilitylink = value
    
    @property
    def newVulnerabilityName(self):
        """ Get newVulnerabilityName value.

          Notes:
              NewVulnerabilityName is the name of the new vulnerability

              
        """
        return self._newvulnerabilityname

    @newVulnerabilityName.setter
    def newVulnerabilityName(self, value):
        """ Set newVulnerabilityName value.

          Notes:
              NewVulnerabilityName is the name of the new vulnerability

              
        """
        self._newvulnerabilityname = value
    
    @property
    def nextPage(self):
        """ Get nextPage value.

          Notes:
              NextPage is the next page number

              
        """
        return self._nextpage

    @nextPage.setter
    def nextPage(self, value):
        """ Set nextPage value.

          Notes:
              NextPage is the next page number

              
        """
        self._nextpage = value
    
    @property
    def notification(self):
        """ Get notification value.

          Notes:
              Notification is the name of the notification sent by Clair using the webhook

              
        """
        return self._notification

    @notification.setter
    def notification(self, value):
        """ Set notification value.

          Notes:
              Notification is the name of the notification sent by Clair using the webhook

              
        """
        self._notification = value
    
    @property
    def notified(self):
        """ Get notified value.

          Notes:
              Notified is the time when the notification was sent

              
        """
        return self._notified

    @notified.setter
    def notified(self, value):
        """ Set notified value.

          Notes:
              Notified is the time when the notification was sent

              
        """
        self._notified = value
    
    @property
    def oldVulnerabilityLink(self):
        """ Get oldVulnerabilityLink value.

          Notes:
              oldVulnerabilityLink is the link that point to the old vulnerability

              
        """
        return self._oldvulnerabilitylink

    @oldVulnerabilityLink.setter
    def oldVulnerabilityLink(self, value):
        """ Set oldVulnerabilityLink value.

          Notes:
              oldVulnerabilityLink is the link that point to the old vulnerability

              
        """
        self._oldvulnerabilitylink = value
    
    @property
    def oldVulnerabilityName(self):
        """ Get oldVulnerabilityName value.

          Notes:
              oldVulnerabilityName is the name of the old vulnerability

              
        """
        return self._oldvulnerabilityname

    @oldVulnerabilityName.setter
    def oldVulnerabilityName(self, value):
        """ Set oldVulnerabilityName value.

          Notes:
              oldVulnerabilityName is the name of the old vulnerability

              
        """
        self._oldvulnerabilityname = value
    
    @property
    def page(self):
        """ Get page value.

          Notes:
              Page is the page number

              
        """
        return self._page

    @page.setter
    def page(self, value):
        """ Set page value.

          Notes:
              Page is the page number

              
        """
        self._page = value
    
    @property
    def parentID(self):
        """ Get parentID value.

          Notes:
              ParentID is the ID of the parent, if any,

              
        """
        return self._parentid

    @parentID.setter
    def parentID(self, value):
        """ Set parentID value.

          Notes:
              ParentID is the ID of the parent, if any,

              
        """
        self._parentid = value
    
    @property
    def parentType(self):
        """ Get parentType value.

          Notes:
              ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.

              
        """
        return self._parenttype

    @parentType.setter
    def parentType(self, value):
        """ Set parentType value.

          Notes:
              ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.

              
        """
        self._parenttype = value
    
    @property
    def status(self):
        """ Get status value.

          Notes:
              Status of an entity

              
        """
        return self._status

    @status.setter
    def status(self, value):
        """ Set status value.

          Notes:
              Status of an entity

              
        """
        self._status = value
    
    @property
    def updatedAt(self):
        """ Get updatedAt value.

          Notes:
              UpdatedAt is the time at which an entity was updated.

              
        """
        return self._updatedat

    @updatedAt.setter
    def updatedAt(self, value):
        """ Set updatedAt value.

          Notes:
              UpdatedAt is the time at which an entity was updated.

              
        """
        self._updatedat = value
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        if len(errors) > 0:
            return errors

        return None

    # clairnotificationIdentity represents the Identity of the object
clairnotificationIdentity = {"name": "clairnotification", "category": "clairnotifications", "constructor": ClairNotification}