## Added the variable for cloud_init data to be passed in
variable "servername" {
  description = "Name of the server"
}

variable "powerinstanceid" {
  description = "Power Instance associated with the account"

  #default="49fba6c9-23f8-40bc-9899-aca322ee7d5b"
  default="59b8fc31-1e13-4f5e-a5dc-c8a999f59647"
}

variable "memory" {
  description = "Memory Of the Power VM Instance"
  default     = "4"
}

variable "processors" {
  description = "Processor Count on the server"
  default     = "1"
}

variable "proctype" {
  description = "Processor Type for the LPAR - shared/dedicated"
  default     = "shared"
}

variable "sshkeyname" {
  description = "Key Name to be passed"
  default     = "suraj-key"
}

variable "volumename" {
  description = "Volume Name to be created"
}

variable "volumesize" {
  description = "Volume Size to be created"
  default     = "40"
}

variable "volumetype" {
  description = "Type of volume to be created - ssd/shared"
  default     = "ssd"
}

variable "shareable" {
  description = "Should the volume be shared or not true/false"
  default     = "true"
}

variable "networks" {
  default = ["Zone1-CFN", "Zone1-IMZ", "Zone1-IBR"]
}

variable "systemtype" {
  description = "Systemtype of the server"
  default     = "e880"
}

variable "migratable" {
  description = "Server can be migrated"
  default     = "false"
}

variable "imagename" {
  description = "Name of the image"
  default     = "7200-04-04"
}

variable "replicationpolicy" {
  description = "Replication Policy of the vm"
  default     = "none"
}

variable "replicants" {
  description = "Number of replicants"
  default     = 1
}

variable "replicant_naming_scheme"
{
description="How to name the created vms"
default="suffix"
}

variable "cloud_init_data" {
  description = "Data to be passed to the instance via cloud init - Must be base64 encoded string"
  default     = "I2Nsb3VkLWNvbmZpZwpzc2hfYXV0aG9yaXplZF9rZXlzOgogLSBmcm9tPSIqLioiLGNvbW1hbmQ9In4vLnNzaC9zc2hkX2NtZF9sb2dnZXIgIFVTOkY6OkFuc2libGUrVXNlcitLZXkiIHNzaC1yc2EgQUFBQUIzTnphQzF5YzJFQUFBQURBUUFCQUFBQ0FRREJKSkptVHEvWFlkSHV3TkxZdXhBbnk2QjU5eFVWcVkvYnZ4OWtOVWhIWnl1S2hOa3RhVlQ2TFlnWmVVV2FFd2NXVTlacWZhRm5qOXZpVDVPNUxPVXRwL1F1T2I0L1Fqem00eWhhU1VST2VPK2FlUk1Kd2JKZEdPRURJRkhaUzVoYmphMzhnb2VoRzkxNjdBUmhJeHJVYytiOUhURXVBNzdReHVjRGVzbnZWblpWTG1kV2xRUDBlNHhjZG5WK0s1bmdCaEN4WTk5aVZKdFBSU1ZVTWZXVURMeFJOSnJtOThrKzA3WG1JL29pc0loalhxZzUxU3VRMHJwcitGSDcrZGZGNGRpdWlrVno1WW1jeHRGcFVsajRpNE55LzhiSjVYNWZmcCtYSURWTEQ0cXdFZkcxck1jb1c2cWpGS1BZV0p5NEpkaFQ3aFVSS0RxdG9MN200R1VjN3o1a0FKNUhtUGRtQ2NBV1lOZWVTNEVTMnQxM0cvWXJxS1ZCcXdZSlZPcjNxZ2ZnZmNnelhEbW1UYVo2VnpUb0FqOHBNZDVJbDF3MzhqUTZwU3ZGK1IwUmZ4ODJBdDZmS3d4QmM5Q2NudVl3dU9WZDdEdW1QVXd2VGlmcFZsOTNsSElSWU9xcFNNVUxPbi9MK2FXTStTNXdCYzg4VXNsSHRtYzhRZ1lTa0k1SVVTSE13WXpMbkZuWlFnNHBGazhvQXhLWFVUUXlTRjlrU2VRQTUvT0lDd21QVUc4dmZ1eThLQm92WnloaEFZUmR5YnVoTHhRQnRlQkN4RjAwZDRMeHRwUitEelFicmE5U3plRFc0S3d1bmNQbHFFMnlST0N6YTdsaEJ2c2RPWWhOaW9IenNUY1gzdmtNeGpZbU1sT0V4ek5IL0pHRFZiYmk2Y0tNNndTNlRRPT0gISFVUy9GLy9JQk0vQW5zaWJsZSBVc2VyIEtleSEhIFVzZWQgZm9yIEFuc2libGUgQXV0b21hdGlvbgogLSBmcm9tPSIqLioiLGNvbW1hbmQ9In4vLnNzaC9zc2hkX2NtZF9sb2dnZXIgIFVTOkY6OkNBTStBdXRvbWF0aW9uK1VzZXIrS2V5IiBzc2gtcnNhIEFBQUFCM056YUMxeWMyRUFBQUFEQVFBQkFBQUNBUURqb2VVNXFUdk9jSDlTYkVKdVVDZ1VZajA1SWdDUWJocm4xTGpoSnZRQktyYWw0RWRXbitsTzIwSHlGMWlaVDBNV05COGVmcDk3NUhoSjQ1WTlzazFGUHg0V2xPejhqamw1Zzc5TnQwV2x0MEFZcWlGcWY0WUR5bnJSdkhPSi8wQXJtUWtGZmtDeFExSzA0aVlEVnAwb3Z1ajA2Mlg5RWl1bTd4UmRlcDd3dXVwVlFtOGVDd2pFcnY5ZUlVK2dvZG1mNjJodndyamsyRVB2YmUxcU5OcEgwRUgrUzZpVGF6RzZlY1JFbmtwRTFKNS8wZVFleDY5TGkvWmFsOFc0bEZyeW96akZrL04yT3oxT1d3Z213bmJpVDdUem9lbkFadzBNejdFNHBGMGNvZGNuZXA2VUpZdmFuNFlBNWsvWVd2WE9nUXh5L1A0MGt0QkFSeWhXa0szM0hJNk1zM2JXQzhPYnBHYlVrV0RDeDdBMmdkOHBKUDYvSTY2M2Y3dmU2RGZNd2NONUZhdTdtM1YvR2FQemVqWU9teEc1Mk5XdzdZVlpSR3VwL1RoWEJsZFV4YmhSVFNJVzVoWld0ZlZrNUlUQndnMHI2MDVuSkptdkp2Mk14eTVwbmdydWN4UTVxVjY4OWtzMVVRQlJodTE0MkdZdHl6amNPSllPQ0F2SE9Eckdyek5Lc1BtamxCQXVrTTlFbGJBQkJ1K2J6akRrbHV5aVVtV3ZzVUpsRXNnb0ZtVmd0YVcwT1pLdEp2MWgzMGVRaDBNSEZyQW9LejRHa0Z6T2VST29UNHVVall6dHNYaUx4ZUEzcFpDTlN0VUxuSGZ2M3MvT0xJcEdvbEt6LzIvUXkycGZOOTRrdEd5LzZNclNnbUtWSzJPRlZyMlRMeGRrTFdZSVM2ZmxmUT09ICEhVVMvRi8vSUJNL0Nsb3VkQXV0b21vYXRpb24gVXNlciBLZXkhISBVc2VkIGZvciBDbG91ZEF1dG9tYXRpb25NYW5hZ2VyIEF1dG9tYXRpb24KcnVuY21kOgogLSBbICdjaGRldicsICctbCcsICdpbmV0MCcsICctYXJvdXRlPW5ldCwtaG9wY291bnQsMSwtbmV0bWFzaywyNTUuMjU1LjI1NS4wLDEwLjI0MC4yLjAse3suSU1aR2F0ZXdheUlQfX0nIF0KIC0gWyAnY2hkZXYnLCAnLWwnLCAnaW5ldDAnLCAnLWFyb3V0ZT1uZXQsLWhvcGNvdW50LDEsLW5ldG1hc2ssMjU1LjI1NS4yNTUuMCwxMC4yNDAuNjQuMCx7ey5JTVpHYXRld2F5SVB9fScgXQogLSBbICdjaGRldicsICctbCcsICdpbmV0MCcsICctYXJvdXRlPW5ldCwtaG9wY291bnQsMSwtbmV0bWFzaywyNTUuMjU1LjI1NS4wLDEwLjI0MC4xMjkuMCx7ey5JTVpHYXRld2F5SVB9fScgXQogLSBbICdjaGRldicsICctbCcsICdpbmV0MCcsICctYXJvdXRlPW5ldCwtaG9wY291bnQsMSwtbmV0bWFzaywyNTUuMjU1LjI1NS4yNDAsMTY5LjU1LjI4LjY0LHt7LklNWkdhdGV3YXlJUH19JyBdCiAtIFsgJ2NoZGV2JywgJy1sJywgJ2luZXQwJywgJy1hcm91dGU9bmV0LC1ob3Bjb3VudCwxLC1uZXRtYXNrLDI1NS4yNTUuMjU1LjI0MCwxNjkuNTUuMTYuMTI4LHt7LklNWkdhdGV3YXlJUH19JyBdCiAtIFsgJ2NoZGV2JywgJy1sJywgJ2luZXQwJywgJy1hcm91dGU9bmV0LC1ob3Bjb3VudCwxLC1uZXRtYXNrLDI1NS4yNTUuMjU1LjI0MCwxMDAuNjQuOC4zMix7ey5JQlJHYXRld2F5SVB9fScgXQogLSBbICdjaGRldicsICctbCcsICdpbmV0MCcsICctYXJvdXRlPW5ldCwtaG9wY291bnQsMSwtbmV0bWFzaywyNTUuMjU1LjI1NS4yMjQsMTE5LjgxLjI4LjMyLHt7LklNWkdhdGV3YXlJUH19JyBdCiAtIFsgJ2NoZGV2JywgJy1sJywgJ2luZXQwJywgJy1hcm91dGU9bmV0LC1ob3Bjb3VudCwxLC1uZXRtYXNrLDI1NS4yNTUuMjU1LjIyNCwxNTkuOC4xODcuMTI4LHt7LklNWkdhdGV3YXlJUH19JyBdCiAtIFsgJ2NoZGV2JywgJy1sJywgJ2luZXQwJywgJy1hcm91dGU9bmV0LC1ob3Bjb3VudCwxLC1uZXRtYXNrLDI1NS4yNTUuMjU1LjIyNCwxNTkuMTIyLjExNS4zMix7ey5JTVpHYXRld2F5SVB9fScgXQogLSBbICdjaGRldicsICctbCcsICdpbmV0MCcsICctYXJvdXRlPW5ldCwtaG9wY291bnQsMSwtbmV0bWFzaywyNTUuMjU1LjI1NS4yMjQsMTYxLjIwMi4zNi4zMix7ey5JTVpHYXRld2F5SVB9fScgXQogLSBbICdjaGRldicsICctbCcsICdpbmV0MCcsICctYXJvdXRlPW5ldCwtaG9wY291bnQsMSwtbmV0bWFzaywyNTUuMjU1LjI1NS4yMjQsMTY5LjU0LjY5Ljk2LHt7LklNWkdhdGV3YXlJUH19JyBdCiAtIFsgJ2NoZGV2JywgJy1sJywgJ2luZXQwJywgJy1hcm91dGU9bmV0LC1ob3Bjb3VudCwxLC1uZXRtYXNrLDI1NS4yNTUuMjU1LjIyNCwxNjkuNTUuMjguMzIse3suSU1aR2F0ZXdheUlQfX0nIF0KIC0gWyAnY2hkZXYnLCAnLWwnLCAnaW5ldDAnLCAnLWFyb3V0ZT1uZXQsLWhvcGNvdW50LDEsLW5ldG1hc2ssMjU1LjI1NS4yNTUuMjI0LDE2OS41NS4xNDIuMjI0LHt7LklNWkdhdGV3YXlJUH19JyBdCiAtIFsgJ2NoZGV2JywgJy1sJywgJ2luZXQwJywgJy1hcm91dGU9bmV0LC1ob3Bjb3VudCwxLC1uZXRtYXNrLDI1NS4yNTUuMjU1LjIyNCwxNjkuNTUuNzkuMjI0LHt7LklNWkdhdGV3YXlJUH19JyBdCiAtIFsgJ2NoZGV2JywgJy1sJywgJ2luZXQwJywgJy1hcm91dGU9bmV0LC1ob3Bjb3VudCwxLC1uZXRtYXNrLDI1NS4yNTUuMjU1LjIyNCwxNjkuNTUuMTkyLjk2LHt7LklNWkdhdGV3YXlJUH19JyBdCiAtIFsgJ2NoZGV2JywgJy1sJywgJ2luZXQwJywgJy1hcm91dGU9bmV0LC1ob3Bjb3VudCwxLC1uZXRtYXNrLDI1NS4yNTUuMjU1LjIyNCwxNjkuNTUuMjU0LjY0LHt7LklNWkdhdGV3YXlJUH19JyBdCiAtIFsgJ2NoZGV2JywgJy1sJywgJ2luZXQwJywgJy1hcm91dGU9bmV0LC1ob3Bjb3VudCwxLC1uZXRtYXNrLDI1NS4yNTUuMjU1LjE5MiwxNjkuNjAuMTM2LjEyOCx7ey5JTVpHYXRld2F5SVB9fScgXQogLSBbICdjaGRldicsICctbCcsICdpbmV0MCcsICctYXJvdXRlPW5ldCwtaG9wY291bnQsMSwtbmV0bWFzaywyNTUuMjU1LjI1NS4xOTIsMTY5LjYwLjEzNi4xOTIse3suSU1aR2F0ZXdheUlQfX0nIF0KIC0gWyAnY2hkZXYnLCAnLWwnLCAnaW5ldDAnLCAnLWFyb3V0ZT1uZXQsLWhvcGNvdW50LDEsLW5ldG1hc2ssMjU1LjI1NS4yNTUuMTkyLDE2OS42Mi4yMTIuNjQse3suSU1aR2F0ZXdheUlQfX0nIF0KIC0gWyAnY2hkZXYnLCAnLWwnLCAnaW5ldDAnLCAnLWFyb3V0ZT1uZXQsLWhvcGNvdW50LDEsLW5ldG1hc2ssMjU1LjI1NS4yNTIuMCwxNDYuODkuMTQwLjAse3suSU1aR2F0ZXdheUlQfX0nIF0KIC0gWyAnY2hkZXYnLCAnLWwnLCAnaW5ldDAnLCAnLWFyb3V0ZT1uZXQsLWhvcGNvdW50LDEsLW5ldG1hc2ssMjU1LjI1NS4yNDguMCwxNDYuODkuMTY4LjAse3suSU1aR2F0ZXdheUlQfX0nIF0KIC0gWyAnY2hkZXYnLCAnLWwnLCAnaW5ldDAnLCAnLWFyb3V0ZT1uZXQsLWhvcGNvdW50LDEsLW5ldG1hc2ssMjU1LjI1NS4yNTQuMCwxNTguODcuNDQuMCx7ey5JTVpHYXRld2F5SVB9fScgXQogLSBbICdjaGRldicsICctbCcsICdpbmV0MCcsICctYXJvdXRlPW5ldCwtaG9wY291bnQsMSwtbmV0bWFzaywyNTUuMjU1LjI1NC4wLDE1OC44Ny40Ni4wLHt7LklNWkdhdGV3YXlJUH19JyBdCiAtIFsgJ2NoZGV2JywgJy1sJywgJ2luZXQwJywgJy1hcm91dGU9bmV0LC1ob3Bjb3VudCwxLC1uZXRtYXNrLDI1NS4yNTUuMjU1LjE5Mix7ey5UU01TZXJ2aWNlSVB9fSx7ey5JQlJHYXRld2F5SVB9fScgXQogLSBbICdzaCcsICcteGMnLCAiZWNobyAnbmFtZXNlcnZlciAxNDYuODkuMTQwLjEyJyA+PiAvZXRjL3Jlc29sdi5jb25mIiBdCiAtIFsgJ3NoJywgJy14YycsICJlY2hvICduYW1lc2VydmVyIDE0Ni44OS4xNDAuMTMnID4+IC9ldGMvcmVzb2x2LmNvbmYiIF0KIC0gWyAncm0nLCAnL2V0Yy9udHAuY29uZiddCiAtIFsgJ3NoJywgJy14YycsICJlY2hvICdkcmlmdGZpbGUgL2V0Yy9udHAuZHJpZnQnID4+IC9ldGMvbnRwLmNvbmYiIF0KIC0gWyAnc2gnLCAnLXhjJywgImVjaG8gJ3RyYWNlZmlsZSAvZXRjL250cC5kcmlmdCcgPj4gL2V0Yy9udHAuY29uZiIgXQogLSBbICdzaCcsICcteGMnLCAiZWNobyAnc2VydmVyIDE2OS42MC4xMzYuMTkzJyA+PiAvZXRjL250cC5jb25mIiBdCiAtIFsgJ2NodmcnLCAnLWcnLCAncm9vdHZnJyBdCiAtIFsgJ2NobHYnLCAnLXgnLCAnMTAwMCcsICdoZDYnIF0KIC0gWyAnY2hmcycsICctYScsICdzaXplPTVHJywgJy9ob21lJyBdCiAtIFsgJ2NoZnMnLCAnLWEnLCAnc2l6ZT0xMEcnLCAnL29wdCcgXQogLSBbICdjaGZzJywgJy1hJywgJ3NpemU9MTBHJywgJy91c3InIF0KIC0gWyAnY2hmcycsICctYScsICdzaXplPTEwRycsICcvdmFyJyBdCiAtIFsgJ2NoZnMnLCAnLWEnLCAnc2l6ZT0xMEcnLCAnL3RtcCcgXQogLSBbICdta2x2JywgJy15JywgJ2xvZ2x2JywgJy10JywgJ2pmczInLCAncm9vdHZnJywgJzQwMCcgXQogLSBbICdjcmZzJywgJy12JywgJ2pmczInLCAnLWQnLCAnbG9nbHYnLCAnLW0nLCAnL3Zhci9sb2cnLCAnLUEnLCAneWVzJywgJy1wJywgJ3J3JyAsJy1hJywgJ3NpemU9MTBHJyBdCiAtIFsgJ21vdW50JywgJy92YXIvbG9nJyBdCmZpbmFsX21lc3NhZ2U6ICJUaGUgc3lzdGVtIGlzIGZpbmFsbHkgdXAiCm91dHB1dCA6IHsgYWxsIDogJ3wgdGVlIC1hIC92YXIvbG9nL2Nsb3VkLWluaXQtb3V0cHV0LmxvZycgfQo="
}
