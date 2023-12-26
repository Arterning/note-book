


## JWT Strategy
```js
import { Injectable, UnauthorizedException } from '@nestjs/common';
import { PassportStrategy } from '@nestjs/passport';
import { ExtractJwt, Strategy } from 'passport-jwt';
import { jwtConstants } from '../constants';
import { UserService } from '../../user/user.service';

@Injectable()
export class JwtStrategy extends PassportStrategy(Strategy) {
  constructor(private userService: UserService) {
    super({
      jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken(),
      ignoreExpiration: false,
      secretOrKey: jwtConstants.secret,
    });
  }


  /**
  the return user object will add to request
  then we can get user like this : request.user
  validate accept the paylod object(extract from token string)
  and get the user id(payload.sub) to find the user exist or not
  **/
  async validate(payload: any) {
    const existUser = this.userService.findOne(payload.sub);

    if (!existUser) {
      throw new UnauthorizedException();
    }

    return { ...payload, id: payload.sub };
  }
}

```

where is the payload constructed, in the login mehod

```js
  async login(user: User) {
    const { password, ...restUser } = user;

    const payload = { ...restUser, sub: user.id };

    return {
      token: this.jwtService.sign(payload),
      user: restUser,
      expiresIn: jwtConstants.expiresIn,
    };
  }
```


## Loal Strategy

```js
import { PassportStrategy } from '@nestjs/passport';
import { Strategy } from 'passport-local';
import { AuthService } from '../auth.service';
import { User } from '../../user/entities/user.entity';
import { Injectable, UnauthorizedException } from '@nestjs/common';
import { ContextIdFactory, ModuleRef } from '@nestjs/core';
import { ReportLogger } from '../../log/ReportLogger';

@Injectable()
export class LocalStrategy extends PassportStrategy(Strategy) {
  constructor(
    private moduleRef: ModuleRef,
    private reportLogger: ReportLogger,
  ) {
    super({ passReqToCallback: true });
    this.reportLogger.setContext('LocalStrategy');
  }

  async validate(
    request: Request,
    username: string,
    password: string,
  ): Promise<Omit<User, 'password'>> {
    const contextId = ContextIdFactory.getByRequest(request);

    // 现在 authService 是一个 request-scoped provider
    const authService = await this.moduleRef.resolve(AuthService, contextId);

    const user = await authService.validateUser(username, password);

    if (!user) {
      this.reportLogger.error('无法登录，SB');
      throw new UnauthorizedException();
    }

    return user;
  }
}

```



## WsJWT Strategy

```js
import { Injectable } from '@nestjs/common';
import { PassportStrategy } from '@nestjs/passport';
import { Strategy } from 'passport-jwt';
import { jwtConstants } from '../constants';
import { UserService } from '../../user/user.service';
import { WsException } from '@nestjs/websockets';

@Injectable()
export class WsJwtStrategy extends PassportStrategy(Strategy, 'ws-jwt') {
  constructor(private userService: UserService) {
    super({
      jwtFromRequest: (req) => {
        const { authorization } = req.handshake.headers;
        if (!authorization) {
          return null;
        }

        const [, token] = authorization.split(' ');

        return token;
      },
      ignoreExpiration: false,
      secretOrKey: jwtConstants.secret,
    });
  }

  async validate(payload: any) {
    const existUser = this.userService.findOne(payload.sub);

    if (!existUser) {
      throw new WsException('无法通过验证');
    }

    return { ...payload, id: payload.sub };
  }
}

```

compare with the JWTStrategy, just a small diff, let's see here the logic to get jwt token method `jwtFromRequest`, it just get the header named `authorization` and get split it then get second string.
```js
constructor(private userService: UserService) {
    super({
      jwtFromRequest: (req) => {
        const { authorization } = req.handshake.headers;
        if (!authorization) {
          return null;
        }

        const [, token] = authorization.split(' ');

        return token;
      },
      ignoreExpiration: false,
      secretOrKey: jwtConstants.secret,
    });
  }
```


so we can just add http head like this

```http
authorization: Bear {token}
authorization: AnyName {token}
```