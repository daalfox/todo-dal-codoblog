# Figuring out the API
This is the first point where I would stage my changes. This helps me think
about the design upfront, so that I don't have to think about it when I'm
implementing the feature.

The tests are not passing at this stage. That's what we want and thats why we
are not committing at this point. We make our first commit when we get the tests
to pass.

You probably noticed that I'm testing `Create` and `Get` functionality together,
because neither one of them are meaningful without the other one. They are tied
together.

You might argue that we can directly talk to database and create a record to
test the `Get` method. Or we can directly talk to database after we used the
`Create` method to confirm that `Create` succeeded. But that direct call is
exactly what you do in the `Get` and `Create` methods, so why would we repeat
ourselves unnecessarily? Yes, we are testing 2 methods together, but that's
only because they can't really be separated and that's fine! Don't make it too
hard on yourself, there is no such rule that you have to only test 1 feature in
your tests.
