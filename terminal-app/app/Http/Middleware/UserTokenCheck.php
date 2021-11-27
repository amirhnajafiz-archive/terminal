<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Redis;

class UserTokenCheck
{
    /**
     * Handle an incoming request.
     *
     * @param Request $request
     * @param Closure $next
     * @return mixed
     */
    public function handle(Request $request, Closure $next)
    {
        if (Redis::exists($request->get('name'))) { // Check in redis first
            if (Redis::get($request->get('name')) == $request->get('token')) {
                return $next($request);
            } else {
                abort(403);
            }
        } else {
            abort(404);
        }

        return back();
    }
}
