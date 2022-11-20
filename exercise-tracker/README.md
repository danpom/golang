# Exercise Tracker
Basically a CRUD demo app. This should track a users exercises as they perform them including reps and sets, failures etc.

### Rough plan
1. Can start with a gui and basic functionality e.g.just log workout no users or sign ins or anything. A workout will be a struct of a slice of exercise. Exercise can be struct with two elements reps and sets.
	Drop down menu of pre-created exercises and option to create exercise (should be per user i.e. if one user creates a new exercise other users should not see it however we may want to avoid recreating the same excercise multiple times so should work on visibility.

2. Add users and sign in functionality 
3. Add plate calculator func. i.e. given the plate denominations available and a workout weight calculate which plates should be added to the bar to get to that weight.
4. pre-configured workout routines that users can select from and customise 

