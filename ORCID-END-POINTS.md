
# ORCID API End Points

+ /orcid-profile/ returns a full profile including both Bio and Activities with API
    + this is what you would use for system integration with the service
+ /orcid-bio/ returns a profile object without the details of activity
    + this is what you want if you want to render output to a web page like a faculty bio
+ /orcid-works/ returns a profile object with only the works section 
    + this is what you want of you are generating a pubs list from an ORCID id
