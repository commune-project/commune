CREATE OR REPLACE FUNCTION public.timestamp_id_secure_random_hex()
  RETURNS text LANGUAGE plpgsql IMMUTABLE PARALLEL SAFE
AS $BODY$
BEGIN
  RETURN SUBSTRING(md5('REPLACE_ME'), 1, 16);
END
$BODY$;;

CREATE OR REPLACE FUNCTION "timestamp_id_fake"(
	table_name text,
  fake_table_name text
  )
    RETURNS bigint
    LANGUAGE 'plpgsql'

    COST 100
    VOLATILE 
AS $BODY$
DECLARE
    time_part bigint;
    sequence_base bigint;
    tail bigint;
  BEGIN
    time_part := (
      -- Get the time in milliseconds
      ((date_part('epoch', now()) * 1000))::bigint
      -- And shift it over two bytes
      << 16);

    sequence_base := (
      'x' ||
      -- Take the first two bytes (four hex characters)
      substr(
        -- Of the MD5 hash of the data we documented
        md5(fake_table_name ||
          timestamp_id_secure_random_hex() ||
          time_part::text
        ),
        1, 4
      )
    -- And turn it into a bigint
    )::bit(16)::bigint;

    -- Finally, add our sequence number to our base, and chop
    -- it to the last two bytes
    tail := (
      (sequence_base + nextval(table_name || '_id_seq'))
      & 65535);

    -- Return the time part and the sequence part. OR appears
    -- faster here than addition, but they're equivalent:
    -- time_part has no trailing two bytes, and tail is only
    -- the last two bytes.
    RETURN time_part | tail;
  END
$BODY$;;

CREATE OR REPLACE FUNCTION "timestamp_id"(
	table_name text
  )
    RETURNS bigint
    LANGUAGE 'plpgsql'

    COST 100
    VOLATILE 
AS $BODY$
BEGIN
  RETURN timestamp_id_fake(table_name, table_name);
END
$BODY$;;

ALTER TABLE "posts" ALTER "id" SET NOT NULL;
ALTER TABLE "posts" ALTER "id" SET DEFAULT timestamp_id_fake('posts'::text, 'statuses'::text);
