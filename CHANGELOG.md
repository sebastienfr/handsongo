# Changelog

## v0.0.4 [16/08/09]
- 16/08/09 chore(doc): update README (SFR)
           chore(deploy): switch to smaller image alpine
           refact(all): use anonymous attribute for router
           refact(all): main action return errors

## v0.0.3 [16/07/13]
- 16/07/07 chore(vendor): update codegansta vendor to urfave (SFR)
           chore(docker): update Go version to 1.6 in Docker deployment
           chore(build): add travis build support

## v0.0.2 [16/06/06]
- 16/06/05 refact(dao): change dao Upsert signature to remove mgo dependency (SFR)
           refact(model): fix model JSON annotation for age
           fix(test): fix statistics test for expected request count
           refact(web): refactor web server in a separate file and main entry point
           feat(test): add full web server test
           chore(test): add benchmarking target and instruction
- 16/06/01 chore(test): add go routine to statistics test (SFR)

## v0.0.1 [16/05/31]
- 16/05/25 chore(test): fix tests to run from all OS (SFR/RLE)
- 16/05/19 feat(all): finalize web server, tests and test script (SFR)
           feat(web): add web server, middleware and router and get implementation without range
           refact(dao): fix panic on wrong object id format
           chore(vendor): update vendors with web dependency
           chore(etc): add scripts to query API
- 16/05/11 feat(dao): add model of spirits, mongo dao, factory and update vendors (SFR)
- 16/05/10 feat(all): project start, add command line, vendors, makefile, docker, logger, tests and statistics (SFR)
