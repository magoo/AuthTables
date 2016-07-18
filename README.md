# AuthTables

AuthTables is a service that assists in the detects when an authentication may have been taken over because of credential theft. If people are stealing your users passwords, AuthTables may be useful.

AuthTables requires zero integration with external feeds of data, risk scores, or machine learning to be useful, though they can certainly be used. Your own authentication data will generate a graph of known locations for a user as they authenticate with known cookies or IP address. Every new login from a previously known IP or Cookie makes this graph stronger over time as it adds new locations for the user.

## Threat

AuthTables is solely focused on the credential theft problem. This is when an attacker has a victim's username and password, but they are not on the victim's host or network. This requires an attacker to authenticate from a different location with a different machine, appearing very different than a normal login. This the most common and most accessible attack that results from large credential dumps and shared passwords.

By being so simple and accessible, simple credential theft and ATO generally makes up for far more than  >50% of the abuse issues related to ATO, while the constellation of other problems (local malware, malicious browser extensions, MITM) usually make up the rest at most companies. The former is fairly simple to defend against with AuthTables, allowing support and engineering attention to be paid to more complicated attacks.

## Opportunity
The attack limitations of a credential thief creates an opportunity for us to build an ever growing graph of known locations a user authenticates from. A credential thief is limited to operating outside of this graph, thus allowing us to treat those authentication with suspicion.

You application may have methods to add locations to this graph, for example:

- Email registrations or link clicks
- Multifactor authentications
- Other risk-based signals based on ML

AuthTables assumes that your authentication service assigns as-static-as-possible cookies or identifiers to your users clients, as it uses this to learn new IP addresses your users connect from.

This allows less friction to the user and greatly reduces the need to prompt for MFA or other out-of-band-verifications.

## Protocol

AuthTables receives JSON POSTs  to `/check` containing UID, IP, and Machine ID.

`{
  "ip":"1.1.1.1",
  "mid":"uniqueidentifier",
  "uid":"magoo"
  }`

AuthTables quickly responds whether this is a known location for the user. If either MID or IP is new, it will add this to their known locations (Response: "OK") which grows their graph. If both are new, there is significant possibility that this account is taken over (Response: "BAD"), and should trigger a multifactor or email confirmation or other way of mitigating risk of ATO for this session. After this challenge, you can `/add` the session to their graph, allowing them to operate in the future without challenges.

## Limitations

- Extra Paranoid users who frequently change hosts and clear cookies (VPN's and Incognito) will frequently appear as credential thiefs
