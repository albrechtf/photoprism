package acl

// Rules specifies granted permissions by Resource and Role.
var Rules = ACL{
	ResourceFiles: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleClient: GrantFullAccess,
	},
	ResourceFolders: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleVisitor: GrantSearchShared,
		RoleClient:  GrantFullAccess,
	},
	ResourceShares: Roles{
		RoleAdmin: GrantFullAccess,
	},
	ResourcePhotos: GrantDefaults,
	ResourceVideos: GrantDefaults,
	ResourceFavorites: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleClient: GrantFullAccess,
	},
	ResourceAlbums: GrantDefaults,
	ResourceMoments: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleVisitor: GrantSearchShared,
		RoleClient:  GrantFullAccess,
	},
	ResourceCalendar: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleVisitor: GrantSearchShared,
		RoleClient:  GrantFullAccess,
	},
	ResourcePeople: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleClient: GrantFullAccess,
	},
	ResourcePlaces: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleVisitor: GrantViewShared,
		RoleClient:  GrantFullAccess,
	},
	ResourceLabels: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleClient: GrantFullAccess,
	},
	ResourceConfig: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleClient:  GrantViewOwn,
		RoleDefault: GrantViewOwn,
	},
	ResourceSettings: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleVisitor: GrantViewOwn,
		RoleClient:  GrantViewUpdateOwn,
	},
	ResourceServices: Roles{
		RoleAdmin: GrantFullAccess,
	},
	ResourcePasscode: Roles{
		RoleAdmin: GrantFullAccess,
	},
	ResourcePassword: Roles{
		RoleAdmin: GrantFullAccess,
	},
	ResourceUsers: Roles{
		RoleAdmin:  GrantAll,
		RoleClient: GrantViewOwn,
	},
	ResourceSessions: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleDefault: GrantOwn,
	},
	ResourceLogs: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleClient: GrantFullAccess,
	},
	ResourceWebDAV: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleClient: GrantFullAccess,
	},
	ResourceMetrics: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleClient: GrantViewAll,
	},
	ResourceFeedback: Roles{
		RoleAdmin: GrantFullAccess,
	},
	ResourceDefault: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleClient: GrantNone,
	},
}
