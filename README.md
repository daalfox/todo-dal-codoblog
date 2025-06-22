# Fixing the tests

This is where we can finally make our first commit. We've compiled the tests
successfully and we provided a migration to initialize our database, so that
the operations on `todo` table don't get thrown into the void. We also handled
errors in this step. It's all about small steps, you can seperate these steps
(fixing tests, adding migrations, handling errors) and stage your code at every
steps to reduce errors, it depends on how much load you can handle at a time.

With the provided migration, the tests compile and pass successfully. Our
feature implementation is now complete.

At this point it's up to you, you can do some house keeping and clean up your
code, or you can start to implement the next feature set.
