package api

type Controllers struct {
	HealthController *HealthController `inject:""`
	StaffController  *StaffController  `inject:""`
}
