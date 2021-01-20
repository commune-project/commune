ALTER TABLE "posts" ADD "boost_to_id" bigint;
ALTER TABLE "posts" ADD CONSTRAINT "fk_posts_boost_to" FOREIGN KEY ("boost_to_id") REFERENCES "posts"("id")